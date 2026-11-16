"""Regressions for consistent Android JDK selection and failure handling."""
import os
from pathlib import Path
import subprocess
import tempfile
import unittest
from unittest.mock import patch

import android_toolchain as toolchain
import build_android
import setup_android


class JavaToolchainTests(unittest.TestCase):
    def setUp(self):
        self.directory = tempfile.TemporaryDirectory()
        self.addCleanup(self.directory.cleanup)
        self.home = Path(self.directory.name) / 'jdk'
        (self.home / 'bin').mkdir(parents=True)
        for name in ('java', 'javac', 'jar', 'keytool'):
            p = self.home / 'bin' / name
            p.write_text('#!/bin/sh\nexit 0\n')
            p.chmod(0o755)
        self.env = {'PATH': str(self.home / 'bin'), 'JAVA_HOME': '/wrong/java'}

    def version(self, args, **kwargs):
        text = 'javac 21.0.1' if Path(args[0]).name == 'javac' else 'openjdk version "21.0.1"'
        return subprocess.CompletedProcess(args, 0, text, '')

    def test_selected_compiler_controls_every_java_tool(self):
        with patch.object(toolchain.subprocess, 'run', side_effect=self.version) as run:
            env = toolchain.java_environment(self.env)
        self.assertEqual(env['JAVA_HOME'], str(self.home))
        self.assertEqual(env['PATH'].split(os.pathsep)[0], str(self.home / 'bin'))
        self.assertEqual([Path(c.args[0][0]).name for c in run.call_args_list], ['java', 'javac'])
        self.assertTrue(all(Path(c.args[0][0]).parent == self.home / 'bin' for c in run.call_args_list))

    def test_rejects_compiler_mismatch_even_when_runtime_matches(self):
        def version(args, **kwargs):
            result = self.version(args, **kwargs)
            if Path(args[0]).name == 'javac':
                result.stdout = 'javac 25.0.1'
            return result
        with patch.object(toolchain.subprocess, 'run', side_effect=version):
            with self.assertRaisesRegex(RuntimeError, 'javac must report major 21'):
                toolchain.java_environment(self.env)

    def test_rejects_failed_or_unparseable_version_commands(self):
        for code, text in [(1, 'openjdk version "21.0.1"'), (0, 'unknown')]:
            with self.subTest(code=code, text=text):
                with patch.object(toolchain.subprocess, 'run', return_value=subprocess.CompletedProcess([], code, text, '')):
                    with self.assertRaises(RuntimeError):
                        toolchain.java_environment(self.env)

    def test_missing_tool_in_selected_home_does_not_mix_jdks(self):
        (self.home / 'bin' / 'jar').unlink()
        with self.assertRaisesRegex(RuntimeError, 'missing executable jar'):
            toolchain.java_environment(self.env)

    def test_local_fallback_only_without_path_compiler(self):
        root = Path(self.directory.name)
        (root / '.tools').mkdir()
        (root / '.tools/jdk').symlink_to(self.home, target_is_directory=True)
        with patch.object(toolchain, 'ROOT', root), patch.object(toolchain.shutil, 'which', return_value=None), patch.object(toolchain.subprocess, 'run', side_effect=self.version):
            env = toolchain.java_environment(self.env)
            self.assertEqual(env['JAVA_HOME'], str(root / '.tools/jdk'))
        # A selected PATH compiler must win even if a local fallback exists.
        with patch.object(toolchain, 'ROOT', root), patch.object(toolchain.subprocess, 'run', side_effect=self.version):
            self.assertEqual(toolchain.java_environment(self.env)['JAVA_HOME'], str(self.home))

    def test_build_and_setup_share_rejection_and_build_checks_before_work(self):
        with patch.object(build_android.Path, 'exists', return_value=True), patch.object(build_android, 'java_environment', side_effect=RuntimeError('bad compiler')):
            self.assertEqual(build_android.unavailable_reason(), 'bad compiler')
            with self.assertRaisesRegex(RuntimeError, 'bad compiler'):
                build_android.build(Path('unused.apk'), 'unused-go', self.env)
        with patch.object(setup_android, 'java_environment', side_effect=RuntimeError('bad compiler')), patch.object(setup_android.shutil, 'which', return_value='/bin/adb'):
            self.assertEqual(setup_android.verify_open_source_toolchain(), ['bad compiler'])


if __name__ == '__main__':
    unittest.main()
