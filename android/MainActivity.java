package dev.cjtrowbridge.stmpck.debug;

import android.app.Activity;
import android.os.Build;
import android.os.Bundle;
import android.view.WindowInsets;
import android.view.View;
import android.widget.FrameLayout;
import dev.cjtrowbridge.stmpck.stmpckmobile.EbitenView;
import dev.cjtrowbridge.stmpck.stmpckmobile.Stmpckmobile;
import go.Seq;

public class MainActivity extends Activity {
    private EbitenView game;

    @Override public void onCreate(Bundle state) {
        super.onCreate(state);
        Seq.setContext(getApplicationContext());
        Stmpckmobile.start(getFilesDir().getAbsolutePath());
        game = new EbitenView(this) {
            @Override protected void onErrorOnGameUpdate(Exception error) {
                throw new RuntimeException("STMPCK game update failed", error);
            }
        };
        FrameLayout root = new FrameLayout(this);
        root.setBackgroundColor(0xff111314);
        root.addView(game, new FrameLayout.LayoutParams(-1, -1));
        root.setOnApplyWindowInsetsListener(new View.OnApplyWindowInsetsListener() {
            @Override public WindowInsets onApplyWindowInsets(View view, WindowInsets insets) {
            if (Build.VERSION.SDK_INT >= 30) {
                android.graphics.Insets bars = insets.getInsets(WindowInsets.Type.systemBars() | WindowInsets.Type.displayCutout());
                view.setPadding(bars.left, bars.top, bars.right, bars.bottom);
            } else {
                view.setPadding(insets.getSystemWindowInsetLeft(), insets.getSystemWindowInsetTop(),
                    insets.getSystemWindowInsetRight(), insets.getSystemWindowInsetBottom());
            }
            return insets;
            }
        });
        setContentView(root);
        game.setKeepScreenOn(true);
        if (Build.VERSION.SDK_INT >= 33) {
            getOnBackInvokedDispatcher().registerOnBackInvokedCallback(
                android.window.OnBackInvokedDispatcher.PRIORITY_DEFAULT,
                new android.window.OnBackInvokedCallback() {
                    @Override public void onBackInvoked() { Stmpckmobile.back(); }
                });
        }
    }

    @Override public void onBackPressed() {
        Stmpckmobile.back();
    }

    @Override protected void onPause() {
        Stmpckmobile.suspend();
        if (game != null) game.suspendGame();
        super.onPause();
    }

    @Override protected void onResume() {
        super.onResume();
        if (game != null) game.resumeGame();
    }
}
