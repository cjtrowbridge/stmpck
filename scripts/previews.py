"""Development-only deterministic sprite capture and bundled-asset validation.

Called by scripts/build.py; uses local Firefox and its Marionette interface.
No third-party Python packages or runtime browser dependencies.
"""
from __future__ import annotations

import base64
import functools
import hashlib
import http.server
import json
import os
from pathlib import Path
import shutil
import socket
import struct
import subprocess
import tempfile
import threading
import time

ROOT = Path(__file__).resolve().parent.parent
ASSETS = ROOT / "internal/preview/assets"


def source_digest():
    paths = sorted({*ROOT.glob("internal/play/*.go"), *ROOT.glob("internal/app/render*.go"), *ROOT.glob("internal/presentation/*.go"),
                    ROOT / "internal/app/game.go", ROOT / "internal/menu/state.go",
                    ROOT / "internal/app/draw.go", ROOT / "internal/app/effects.go",
                    ROOT / "internal/app/preview_capture.go", ROOT / "internal/preview/assets.go",
                    ROOT / "scripts/previews.py", ROOT / "go.mod", ROOT / "go.sum"})
    h = hashlib.sha256()
    for path in paths:
        if path.name.endswith("_test.go"):
            continue
        # Git checkouts may use CRLF on Windows; neither line endings nor
        # host path separators change the captured rules or rendering.
        h.update(path.relative_to(ROOT).as_posix().encode() + b"\0" + path.read_bytes().replace(b"\r\n", b"\n"))
    return h.hexdigest()


def validate():
    manifest = json.loads((ASSETS / "manifest.json").read_text())
    if manifest["version"] != 1 or len(set(manifest["games"])) != len(manifest["games"]):
        raise ValueError("Invalid preview manifest version or game identifiers")
    expected = {f"{game}-{style}-{kind}.png" for game in manifest["games"] for style in range(3) for kind in ("still", "sheet") if style > 0 or kind == "still"}
    if set(manifest["files"]) != expected:
        raise ValueError("Incomplete or extra preview variants")
    if manifest["source_sha256"] != source_digest():
        raise ValueError("Preview sources changed. Run python scripts/build.py --generate-previews and review the captures.")
    for name, entry in manifest["files"].items():
        data = (ASSETS / name).read_bytes()
        if hashlib.sha256(data).hexdigest() != entry["sha256"]:
            raise ValueError(f"Preview checksum mismatch: {name}")
        if data[:8] != b"\x89PNG\r\n\x1a\n" or list(struct.unpack(">II", data[16:24])) != entry["size"]:
            raise ValueError(f"Invalid preview PNG: {name}")
    print(f"✓ Preview manifest: {len(manifest['games'])} games, {len(manifest['files'])} images, {manifest['bytes']:,} bytes", flush=True)
    return manifest


class Browser:
    """Small local-only Marionette client shared by capture and smoke checks."""
    def __init__(self, directory, profile):
        handler = functools.partial(http.server.SimpleHTTPRequestHandler, directory=str(directory))
        self.server = http.server.ThreadingHTTPServer(("127.0.0.1", 0), handler)
        threading.Thread(target=self.server.serve_forever, daemon=True).start()
        self.url = f"http://127.0.0.1:{self.server.server_port}"
        with socket.socket() as probe:
            probe.bind(("127.0.0.1", 0))
            port = probe.getsockname()[1]
        profile.mkdir(parents=True, exist_ok=True)
        (profile / "user.js").write_text(f'user_pref("marionette.port", {port});\n')
        self.process = subprocess.Popen(["firefox", "--headless", "--no-remote", "--marionette", "--profile", str(profile)], stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
        self.sock = None
        self.serial = 0
        for _ in range(120):
            try:
                self.sock = socket.create_connection(("127.0.0.1", port), timeout=1)
                break
            except OSError:
                time.sleep(.25)
        if self.sock is None:
            self.close()
            raise RuntimeError("Firefox Marionette did not start")
        self.sock.settimeout(180)
        self.receive()
        self.call("WebDriver:NewSession", {"capabilities": {}})

    def receive(self):
        header = b""
        while not header.endswith(b":"):
            chunk = self.sock.recv(1)
            if not chunk:
                raise RuntimeError("Firefox connection closed")
            header += chunk
        size = int(header[:-1])
        body = b""
        while len(body) < size:
            chunk = self.sock.recv(size-len(body))
            if not chunk:
                raise RuntimeError("Firefox connection closed")
            body += chunk
        return json.loads(body)

    def call(self, name, args=None):
        self.serial += 1
        payload = json.dumps([0, self.serial, name, args or {}]).encode()
        self.sock.sendall(str(len(payload)).encode()+b":"+payload)
        result = self.receive()
        if result[2]:
            raise RuntimeError(result[2])
        return result[3]

    def js(self, script):
        return self.call("WebDriver:ExecuteScript", {"script": script, "args": [], "newSandbox": False})["value"]

    def close(self):
        if self.sock:
            self.sock.close()
        self.process.terminate()
        try:
            self.process.wait(timeout=5)
        except subprocess.TimeoutExpired:
            self.process.kill()
            self.process.wait()
        self.server.shutdown()
        self.server.server_close()


def generate(go, environment):
    if not shutil.which("firefox"):
        raise RuntimeError("Preview generation needs local Firefox; ordinary builds use checked-in images.")
    tools = ROOT / ".tools"
    tools.mkdir(exist_ok=True)
    with tempfile.TemporaryDirectory(prefix="preview-capture-", dir=tools) as temporary:
        work = Path(temporary)
        env = dict(environment, GOOS="js", GOARCH="wasm", CGO_ENABLED="0")
        subprocess.run([go, "build", "-buildvcs=false", "-tags=previewcapture", "-o", str(work / "app.wasm"), "./cmd/preview"], cwd=ROOT, env=env, check=True)
        goroot = subprocess.check_output([go, "env", "GOROOT"], env=environment, text=True).strip()
        shutil.copyfile(Path(goroot) / "lib/wasm/wasm_exec.js", work / "wasm_exec.js")
        (work / "index.html").write_text('<script src="wasm_exec.js"></script><script>window.captureErrors=[];for(const method of ["error","log"]){const old=console[method];console[method]=(...a)=>{captureErrors.push(a.join(" "));old(...a)}}const go=new Go();go.exit=c=>captureErrors.push("Go exit "+c);WebAssembly.instantiateStreaming(fetch("app.wasm"),go.importObject).then(r=>go.run(r.instance)).catch(e=>captureErrors.push(String(e)));</script>')
        browser = Browser(work, work / "profile")
        try:
            browser.call("WebDriver:Navigate", {"url": browser.url})
            for _ in range(120):
                if browser.js('return typeof capturePreview === "function"'):
                    break
                time.sleep(.25)
            else:
                raise RuntimeError("Capture program did not initialize")
            ids = browser.js("return previewIDs")
            manifest = {"version": 1, "source_sha256": source_digest(), "size": 96, "fps": 12, "columns": 8, "games": ids, "files": {}, "bytes": 0}
            for game, name in enumerate(ids):
                for style in range(3):
                    browser.js(f"return capturePreview({game},{style})")
                    for _ in range(600):
                        errors = browser.js("return captureErrors")
                        if errors:
                            raise RuntimeError(f"Capture failed: {name}/{style}: {errors}")
                        result = browser.js("return previewResult")
                        if result:
                            break
                        time.sleep(.1)
                    else:
                        raise RuntimeError(f"Capture timed out: {name}/{style}")
                    for kind, encoded in result.items():
                        data = base64.b64decode(encoded)
                        filename = f"{name}-{style}-{kind}.png"
                        (work / filename).write_bytes(data)
                        manifest["files"][filename] = {"sha256": hashlib.sha256(data).hexdigest(), "size": list(struct.unpack(">II", data[16:24])), "bytes": len(data)}
                        manifest["bytes"] += len(data)
                print(f"  Captured {name}", flush=True)
            for filename in manifest["files"]:
                os.replace(work / filename, ASSETS / filename)
            (ASSETS / "manifest.json").write_text(json.dumps(manifest, indent=2)+"\n")
        finally:
            browser.close()
    validate()
