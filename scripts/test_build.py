"""Focused build-driver regressions; run with python -m unittest discover -s scripts."""
import subprocess
import tempfile
import unittest
import sys
from pathlib import Path
from unittest.mock import patch

import build
import previews


class WindowsBuildTests(unittest.TestCase):
    def test_ordinary_build_reports_failure_without_verify_flag(self):
        target = build.Target("windows", "amd64", "windows-amd64", ".exe", "latest.exe")
        with (patch.object(sys, "argv", ["build.py"]),
              patch.object(build.previews, "validate"),
              patch.object(build.signal, "signal"),
              patch.object(build, "build_target_matrix", return_value=[target]),
              patch.object(build, "_print_feasibility_report"),
              patch.object(build, "_run_build", return_value=False),
              patch.object(build, "_persist_report", return_value=Path("report.json"))):
            self.assertEqual(build.main(), 1)

    def test_native_windows_build_needs_no_c_compiler(self):
        with tempfile.TemporaryDirectory() as temporary:
            with (patch.object(build, "_host_os", return_value="windows"),
                  patch.object(build, "_host_arch", return_value="amd64"),
                  patch.object(build, "_detect_go_toolchain", return_value="go"),
                  patch.object(build, "_go_environment", return_value={"CGO_ENABLED": "1"}),
                  patch.object(build, "RELEASE_DIR", Path(temporary)),
                  patch.object(build.subprocess, "run", return_value=subprocess.CompletedProcess([], 0, "", "")) as run):
                matrix = build.build_target_matrix()
                self.assertEqual({t.name for t in matrix if t.feasible}, {"windows-amd64", "js-wasm"})
                target = next(t for t in matrix if t.name == "windows-amd64")
                self.assertTrue(build._run_build(target))
                self.assertEqual(target.output_path, Path(temporary) / "windows/amd64/app/latest.exe")
                self.assertEqual(run.call_args.kwargs["env"], {"GOOS": "windows", "GOARCH": "amd64", "CGO_ENABLED": "0"})
                self.assertEqual(run.call_args.args[0][-3:], ["-o", str(target.output_path), "./cmd/app"])

    def test_failed_build_cannot_verify_a_stale_executable(self):
        with tempfile.TemporaryDirectory() as temporary:
            stale = Path(temporary) / "latest.exe"
            stale.write_bytes(b"previous executable")
            target = build.Target("windows", "amd64", "windows-amd64", ".exe", "latest.exe", output_path=stale)
            self.assertFalse(build._verify_artifacts([target]))


class PreviewFingerprintTests(unittest.TestCase):
    def test_checkout_line_endings_are_portable_but_source_edits_invalidate(self):
        with tempfile.TemporaryDirectory() as temporary:
            root = Path(temporary)
            names = ["internal/app/game.go", "internal/menu/state.go", "internal/app/draw.go",
                     "internal/app/effects.go", "internal/app/preview_capture.go", "internal/preview/assets.go",
                     "scripts/previews.py", "go.mod", "go.sum", "internal/play/example.go"]
            for name in names:
                path = root / name
                path.parent.mkdir(parents=True, exist_ok=True)
                path.write_bytes(b"original\ncontent\n")
            with patch.object(previews, "ROOT", root):
                original = previews.source_digest()
                for name in names:
                    (root / name).write_bytes(b"original\r\ncontent\r\n")
                self.assertEqual(previews.source_digest(), original)
                (root / "internal/play/example.go").write_bytes(b"changed rules\r\n")
                self.assertNotEqual(previews.source_digest(), original)


if __name__ == "__main__":
    unittest.main()
