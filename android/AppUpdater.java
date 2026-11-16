package dev.cjtrowbridge.stmpck.debug;

import android.app.Activity;
import android.app.AlertDialog;
import android.content.Intent;
import android.content.DialogInterface;
import android.content.pm.PackageInfo;
import android.content.pm.PackageManager;
import android.net.Uri;
import android.os.Build;
import android.os.SystemClock;
import android.provider.Settings;
import java.io.File;
import java.io.FileInputStream;
import java.io.FilenameFilter;
import java.io.FileOutputStream;
import java.io.InputStream;
import java.net.HttpURLConnection;
import java.net.URL;
import java.security.MessageDigest;
import java.util.Arrays;
import java.util.concurrent.atomic.AtomicBoolean;
import dev.cjtrowbridge.stmpck.stmpckmobile.Stmpckmobile;

// One bounded download per Activity. Only explicit confirmation opens the installer.
// The startup flag is process-local: resuming gameplay never starts another check.
final class AppUpdater {
    private static final String URL_LATEST = "https://raw.githubusercontent.com/cjtrowbridge/stmpck-builds/main/latest.apk";
    private static final AtomicBoolean startupChecked = new AtomicBoolean();
    private final Activity activity;
    private final AtomicBoolean checking = new AtomicBoolean();
    private volatile boolean destroyed, manual;
    private volatile HttpURLConnection connection;
    private Thread worker;
    private boolean resumed, offerPending, permissionPending;
    private File candidate;

    AppUpdater(Activity activity) { this.activity = activity; }

    void startup() {
        if (startupChecked.compareAndSet(false, true)) {
            File[] stale = activity.getCacheDir().listFiles(new FilenameFilter() {
                @Override public boolean accept(File dir, String name) {
                    return name.startsWith("stmpck-update-") && name.endsWith(".part");
                }
            });
            if (stale != null) for (File file : stale) file.delete();
            check(false);
        }
    }

    void resume() {
        resumed = true;
        if (permissionPending) {
            permissionPending = false;
            if (Build.VERSION.SDK_INT < 26 || activity.getPackageManager().canRequestPackageInstalls()) install();
            else status("Install permission not granted. Try again.");
        } else if (offerPending) offer();
    }

    void pause() { resumed = false; }

    void close() {
        destroyed = true;
        HttpURLConnection active = connection;
        if (active != null) active.disconnect();
        if (worker != null) worker.interrupt();
    }

    void check(boolean userRequested) {
        // Both native startup and Go callbacks enter on the Android UI thread.
        if (destroyed) return;
        if (checking.get()) {
            if (userRequested) { manual = true; status("Checking for an update..."); }
            return;
        }
        if (offerPending) { offer(); return; }
        manual = userRequested;
        if (!checking.compareAndSet(false, true)) return;
        if (manual) status("Checking for an update...");
        worker = new Thread(new Runnable() {
          @Override public void run() {
            File temporary = null;
            boolean handedOff = false;
            try {
                File apk = download();
                temporary = apk;
                if (Arrays.equals(hash(apk), hash(new File(activity.getApplicationInfo().sourceDir)))) {
                    apk.delete();
                    if (manual) status("Up to date.");
                } else if (!trustedNewer(apk)) {
                    apk.delete();
                    if (manual) status("No compatible newer update available.");
                } else {
                    handedOff = true;
                    activity.runOnUiThread(new Runnable() {
                      @Override public void run() {
                        if (destroyed) { apk.delete(); return; }
                        File ready = new File(activity.getCacheDir(), "update.apk");
                        if (!apk.renameTo(ready)) {
                            apk.delete();
                            if (manual) status("Unable to prepare update. Try again.");
                            return;
                        }
                        candidate = ready;
                        offerPending = true;
                        status("An update is available.");
                        if (resumed) offer();
                      }
                    });
                }
            } catch (Exception error) {
                if (manual && !destroyed) status("Unable to check. Try again when online.");
            } finally {
                if (!handedOff && temporary != null) temporary.delete();
                // Release on the UI queue after candidate handoff, avoiding a new download race.
                activity.runOnUiThread(new Runnable() {
                    @Override public void run() { checking.set(false); }
                });
            }
          }
        }, "stmpck-update");
        worker.start();
    }

    private File download() throws Exception {
        File partial = File.createTempFile("stmpck-update-", ".part", activity.getCacheDir());
        HttpURLConnection request = (HttpURLConnection) new URL(URL_LATEST).openConnection();
        connection = request;
        boolean complete = false;
        try {
            request.setUseCaches(false);
            request.setInstanceFollowRedirects(false);
            request.setConnectTimeout(15000);
            request.setReadTimeout(15000);
            if (destroyed || request.getResponseCode() != 200) throw new java.io.IOException("update unavailable");
            long deadline = SystemClock.elapsedRealtime() + 120000;
            long total = 0;
            try (InputStream in = request.getInputStream(); FileOutputStream out = new FileOutputStream(partial)) {
                byte[] buffer = new byte[65536];
                int count;
                while ((count = in.read(buffer)) != -1) {
                    total += count;
                    if (destroyed || Thread.currentThread().isInterrupted() || total > 128L * 1024 * 1024
                            || SystemClock.elapsedRealtime() > deadline) throw new java.io.IOException("cancelled or limit exceeded");
                    out.write(buffer, 0, count);
                }
                if (total == 0 || destroyed) throw new java.io.IOException("empty or cancelled");
                out.getFD().sync();
            }
            complete = true;
            return partial;
        } finally {
            request.disconnect();
            connection = null;
            if (!complete) partial.delete();
        }
    }

    @SuppressWarnings("deprecation")
    private boolean trustedNewer(File apk) throws Exception {
        PackageManager manager = activity.getPackageManager();
        PackageInfo current = manager.getPackageInfo(activity.getPackageName(), PackageManager.GET_SIGNATURES);
        PackageInfo next = manager.getPackageArchiveInfo(apk.getPath(), PackageManager.GET_SIGNATURES);
        if (next == null || !current.packageName.equals(next.packageName)) return false;
        long oldCode = Build.VERSION.SDK_INT >= 28 ? current.getLongVersionCode() : current.versionCode;
        long newCode = Build.VERSION.SDK_INT >= 28 ? next.getLongVersionCode() : next.versionCode;
        return newCode > oldCode && current.signatures != null && current.signatures.length > 0
                && next.signatures != null && Arrays.equals(current.signatures, next.signatures);
    }

    private void offer() {
        if (!resumed || destroyed || !offerPending) return;
        offerPending = false;
        new AlertDialog.Builder(activity).setTitle("STMPCK update available")
            .setMessage("Install the newer version? Your saved games will be kept.")
            .setNegativeButton("Later", new DialogInterface.OnClickListener() {
                @Override public void onClick(DialogInterface dialog, int which) {
                    status("Update available. Use Settings to retry.");
                }
            })
            .setPositiveButton("Update", new DialogInterface.OnClickListener() {
              @Override public void onClick(DialogInterface dialog, int which) {
                if (Build.VERSION.SDK_INT >= 26 && !activity.getPackageManager().canRequestPackageInstalls()) {
                    permissionPending = true;
                    try {
                        activity.startActivity(new Intent(Settings.ACTION_MANAGE_UNKNOWN_APP_SOURCES,
                            Uri.parse("package:" + activity.getPackageName())));
                        status("Allow updates from STMPCK to continue.");
                    } catch (Exception error) {
                        permissionPending = false;
                        status("Unable to open install permission settings.");
                    }
                } else install();
              }
            }).show();
    }

    private void install() {
        if (destroyed) return;
        try {
            if (candidate == null || !trustedNewer(candidate)) {
                status("Update verification failed. Try again.");
                return;
            }
            // Flush the latest game snapshot before leaving for Android's installer.
            Stmpckmobile.suspend();
            Intent intent = new Intent(Intent.ACTION_VIEW);
            intent.setDataAndType(Uri.parse("content://" + activity.getPackageName() + ".update/latest.apk"),
                "application/vnd.android.package-archive");
            intent.addFlags(Intent.FLAG_GRANT_READ_URI_PERMISSION);
            activity.startActivity(intent);
            status("Installer opened; confirm in Android.");
        } catch (Exception error) {
            status("Unable to open the Android installer.");
        }
    }

    private void status(String message) {
        if (!destroyed) Stmpckmobile.reportUpdateStatus(message);
    }

    private static byte[] hash(File file) throws Exception {
        MessageDigest digest = MessageDigest.getInstance("SHA-256");
        try (InputStream in = new FileInputStream(file)) {
            byte[] buffer = new byte[65536];
            int count;
            while ((count = in.read(buffer)) != -1) digest.update(buffer, 0, count);
        }
        return digest.digest();
    }
}
