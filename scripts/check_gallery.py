"""Local browser gallery smoke checks and comparable frame-time observations.

Usage: python scripts/check_gallery.py --label before|after [--checks]
Uses the last successfully built WASM artifact; performs no builds or deployment.
"""
import argparse
import base64
import json
from pathlib import Path
import shutil
import statistics
import subprocess
import tempfile
import time

import build
from previews import Browser, ROOT


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--label", required=True)
    parser.add_argument("--checks", action="store_true")
    args = parser.parse_args()
    if not args.label.replace("-", "").isalnum():
        raise ValueError("label must contain only letters, numbers and hyphens")
    output = ROOT / "reports/sprite-previews" / args.label
    output.mkdir(parents=True, exist_ok=True)
    with tempfile.TemporaryDirectory(prefix="gallery-check-", dir=ROOT / ".tools") as temporary:
        work = Path(temporary)
        shutil.copyfile(ROOT / "releases/js/wasm/app/latest.wasm", work / "app.wasm")
        goroot = subprocess.check_output([build._detect_go_toolchain(), "env", "GOROOT"], env=build._go_environment(), text=True).strip()
        shutil.copyfile(Path(goroot) / "lib/wasm/wasm_exec.js", work / "wasm_exec.js")
        (work / "index.html").write_text('<meta name="viewport" content="width=device-width,initial-scale=1"><style>body{margin:0;overflow:hidden}</style><script>localStorage.clear();</script><script src="wasm_exec.js"></script><script>const go=new Go();WebAssembly.instantiateStreaming(fetch("app.wasm"),go.importObject).then(r=>go.run(r.instance));</script>')
        browser = Browser(work, work / "profile")
        try:
            browser.call("WebDriver:SetWindowRect", {"width": 480, "height": 1000})
            browser.call("WebDriver:Navigate", {"url": browser.url})
            time.sleep(5)

            def key(value):
                browser.call("WebDriver:PerformActions", {"actions": [{"type": "key", "id": "keys", "actions": [{"type": "keyDown", "value": value}, {"type": "pause", "duration": 60}, {"type": "keyUp", "value": value}]}]})
                time.sleep(.15)

            def shot(name):
                data = base64.b64decode(browser.call("WebDriver:TakeScreenshot", {"id": None, "full": False, "scroll": False})["value"])
                (output / (name + ".png")).write_bytes(data)
                return data

            def sample():
                browser.js('window.samples=[];window.last=0;window.until=performance.now()+10000;requestAnimationFrame(function f(t){if(window.last)window.samples.push(t-window.last);window.last=t;if(t<window.until)requestAnimationFrame(f);});return true')
                time.sleep(11)
                values = sorted(browser.js("return window.samples"))
                return {"frames": len(values), "median_ms": statistics.median(values), "p95_ms": values[int(len(values)*.95)], "p99_ms": values[int(len(values)*.99)], "max_ms": max(values), "over_50ms": sum(v>50 for v in values)}

            result = {"viewport": browser.js("return [innerWidth,innerHeight]"), "default": sample(), "phone_tested": False}
            shot("gallery-default")
            assert browser.js('return localStorage.getItem("stmpck/games.json")') is None, "previews wrote player saves"
            if args.checks:
                key("\ue012")
                still = shot("still")
                time.sleep(.7)
                assert still == shot("still-later"), "still previews changed"
                key("\ue012")
                wire = shot("wireframe")
                assert wire != still, "wireframe did not change presentation"
                time.sleep(.7)
                assert wire == shot("wireframe-later"), "wireframe moved"
                key("\ue014")
                assert shot("still-return") == still, "default still did not preserve its frame"
                key("\ue014")
                time.sleep(.5)
                assert shot("resumed") != still, "default animation did not resume"
                for _ in range(5): key("\ue014")
                time.sleep(4)
                shot("maximum")
                result["maximum"] = sample()
                for _ in range(17): key("\ue015")
                time.sleep(2)
                shot("scrolled")
                for _ in range(5): key("\ue012")
                key("h")
                settings = shot("settings")
                time.sleep(.5)
                assert settings == shot("settings-later"), "settings changed"
                key("\ue00c")
                assert browser.js('return localStorage.getItem("stmpck/games.json")') is None
                key("\ue006")
                shot("setup")
                assert browser.js('return localStorage.getItem("stmpck/games.json")') is None, "opening setup wrote a save"
                key("\ue006")
                shot("game")
                saved = json.loads(browser.js('return localStorage.getItem("stmpck/games.json")'))
                assert sum(saved["Plays"]) == 1, "starting the game did not record exactly one play"
                key("\ue00c")
                shot("returned")
                result["checks"] = ["still freeze", "wireframe", "same-frame still return", "animation resume", "maximum intensity", "scroll", "settings", "game entry/return", "no preview saves or counts"]
            (output / "checks.json").write_text(json.dumps(result, indent=2)+"\n")
            print(json.dumps(result, indent=2), flush=True)
        finally:
            browser.close()


if __name__ == "__main__":
    main()
