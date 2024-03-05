package main

import (
	"bytes"
	"testing"
)

func TestHash(t *testing.T) {
	t.Run("unsupported hashing algorithm", func(t *testing.T) {
		var (
			args = []string{"hash", "-a", "md4"}
			in   = bytes.NewReader([]byte(""))
			out  = new(bytes.Buffer)
			err  = new(bytes.Buffer)
		)

		exitCode := run(args, in, out, err)

		assertEqual(t, exitCode, 1)
		assertEqual(t, out.String(), "")
		assertEqual(t, err.String(), "Unsupported hashing algorithm\n")
	})

	t.Run("help", func(t *testing.T) {
		helpMsg := `hash program usage:

hash [FLAGS] -a SHA1 [STDIN]
hash [FLAGS] -a SHA1 -- [STDIN]
hash [FLAGS] -a SHA1 [FILE]

Flags:
-a, -algorithm one of {MD5 SHA1 SHA224 SHA256 SHA384 SHA512}
-h, -help print help
-v, -version current version
`
		t.Run("-h", func(t *testing.T) {
			var (
				args = []string{"hash", "-h"}
				in   = bytes.NewReader([]byte(""))
				out  = new(bytes.Buffer)
				err  = new(bytes.Buffer)
			)

			exitCode := run(args, in, out, err)

			assertEqual(t, exitCode, 2)
			assertEqual(t, out.String(), "")
			assertEqual(t, err.String(), helpMsg)
		})

		t.Run("-help", func(t *testing.T) {
			var (
				args = []string{"hash", "-help"}
				in   = bytes.NewReader([]byte(""))
				out  = new(bytes.Buffer)
				err  = new(bytes.Buffer)
			)

			exitCode := run(args, in, out, err)

			assertEqual(t, exitCode, 2)
			assertEqual(t, out.String(), "")
			assertEqual(t, err.String(), helpMsg)
		})

		t.Run("incomplite flag", func(t *testing.T) {
			var (
				args = []string{"hash", "-"}
				in   = bytes.NewReader([]byte(""))
				out  = new(bytes.Buffer)
				err  = new(bytes.Buffer)
			)

			exitCode := run(args, in, out, err)

			assertEqual(t, exitCode, 2)
			assertEqual(t, out.String(), "")
			assertEqual(t, err.String(), helpMsg)
		})

	})

	t.Run("hashing from stdin", func(t *testing.T) {
		t.Run("-a MD5", func(t *testing.T) {
			var (
				args = []string{"hash", "-a", "md5"}
				in   = bytes.NewReader([]byte("Hello World"))
				out  = new(bytes.Buffer)
				err  = new(bytes.Buffer)
			)

			exitCode := run(args, in, out, err)

			assertEqual(t, exitCode, 0)
			assertEqual(t, out.String(), "b10a8db164e0754105b7a99be72e3fe5\n")
			assertEqual(t, err.String(), "")
		})

		t.Run("-algorithm SHA1", func(t *testing.T) {
			var (
				args = []string{"hash", "-algorithm", "sha1", "--"}
				in   = bytes.NewReader([]byte("Hello World"))
				out  = new(bytes.Buffer)
				err  = new(bytes.Buffer)
			)

			exitCode := run(args, in, out, err)

			assertEqual(t, exitCode, 0)
			assertEqual(t, out.String(), "0a4d55a8d778e5022fab701977c5d840bbc486d0\n")
			assertEqual(t, err.String(), "")
		})

		t.Run("-a SHA224", func(t *testing.T) {
			var (
				args = []string{"hash", "-a", "sha224"}
				in   = bytes.NewReader([]byte("Hello World"))
				out  = new(bytes.Buffer)
				err  = new(bytes.Buffer)
			)

			exitCode := run(args, in, out, err)

			assertEqual(t, exitCode, 0)
			assertEqual(t, out.String(), "c4890faffdb0105d991a461e668e276685401b02eab1ef4372795047\n")
			assertEqual(t, err.String(), "")
		})

		t.Run("-a SHA256", func(t *testing.T) {
			var (
				args = []string{"hash", "-a", "sha256", "--"}
				in   = bytes.NewReader([]byte("Hello World"))
				out  = new(bytes.Buffer)
				err  = new(bytes.Buffer)
			)

			exitCode := run(args, in, out, err)

			assertEqual(t, exitCode, 0)
			assertEqual(t, out.String(), "a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e\n")
			assertEqual(t, err.String(), "")
		})

		t.Run("-a SHA384", func(t *testing.T) {
			var (
				args = []string{"hash", "-a", "sha384"}
				in   = bytes.NewReader([]byte("Hello World"))
				out  = new(bytes.Buffer)
				err  = new(bytes.Buffer)
			)

			exitCode := run(args, in, out, err)

			assertEqual(t, exitCode, 0)
			assertEqual(t, out.String(), "99514329186b2f6ae4a1329e7ee6c610a729636335174ac6b740f9028396fcc803d0e93863a7c3d90f86beee782f4f3f\n")
			assertEqual(t, err.String(), "")
		})

		t.Run("-a SHA512", func(t *testing.T) {
			var (
				args = []string{"hash", "-a", "sha512"}
				in   = bytes.NewReader([]byte("Hello World"))
				out  = new(bytes.Buffer)
				err  = new(bytes.Buffer)
			)

			exitCode := run(args, in, out, err)

			assertEqual(t, exitCode, 0)
			assertEqual(t, out.String(), "2c74fd17edafd80e8447b0d46741ee243b7eb74dd2149a0ab1b9246fb30382f27e853d8585719e0e67cbda0daa8f51671064615d645ae27acb15bfb1447f459b\n")
			assertEqual(t, err.String(), "")
		})
	})

	t.Run("hashing from file", func(t *testing.T) {
		t.Run("-a MD5", func(t *testing.T) {
			var (
				args = []string{"hash", "-a", "md5", "./fixtures/example.md"}
				in   = bytes.NewReader(nil)
				out  = new(bytes.Buffer)
				err  = new(bytes.Buffer)
			)

			exitCode := run(args, in, out, err)

			assertEqual(t, exitCode, 0)
			assertEqual(t, out.String(), "b10a8db164e0754105b7a99be72e3fe5\n")
			assertEqual(t, err.String(), "")
		})

		t.Run("-algorithm SHA1", func(t *testing.T) {
			var (
				args = []string{"hash", "-algorithm", "sha1", "./fixtures/example.md"}
				in   = bytes.NewReader(nil)
				out  = new(bytes.Buffer)
				err  = new(bytes.Buffer)
			)

			exitCode := run(args, in, out, err)

			assertEqual(t, exitCode, 0)
			assertEqual(t, out.String(), "0a4d55a8d778e5022fab701977c5d840bbc486d0\n")
			assertEqual(t, err.String(), "")
		})

		t.Run("-a SHA224", func(t *testing.T) {
			var (
				args = []string{"hash", "-a", "sha224", "./fixtures/example.md"}
				in   = bytes.NewReader(nil)
				out  = new(bytes.Buffer)
				err  = new(bytes.Buffer)
			)

			exitCode := run(args, in, out, err)

			assertEqual(t, exitCode, 0)
			assertEqual(t, out.String(), "c4890faffdb0105d991a461e668e276685401b02eab1ef4372795047\n")
			assertEqual(t, err.String(), "")
		})

		t.Run("-a SHA256", func(t *testing.T) {
			var (
				args = []string{"hash", "-a", "sha256", "./fixtures/example.md"}
				in   = bytes.NewReader(nil)
				out  = new(bytes.Buffer)
				err  = new(bytes.Buffer)
			)

			exitCode := run(args, in, out, err)

			assertEqual(t, exitCode, 0)
			assertEqual(t, out.String(), "a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e\n")
			assertEqual(t, err.String(), "")
		})

		t.Run("-a SHA384", func(t *testing.T) {
			var (
				args = []string{"hash", "-a", "sha384", "./fixtures/example.md"}
				in   = bytes.NewReader(nil)
				out  = new(bytes.Buffer)
				err  = new(bytes.Buffer)
			)

			exitCode := run(args, in, out, err)

			assertEqual(t, exitCode, 0)
			assertEqual(t, out.String(), "99514329186b2f6ae4a1329e7ee6c610a729636335174ac6b740f9028396fcc803d0e93863a7c3d90f86beee782f4f3f\n")
			assertEqual(t, err.String(), "")
		})

		t.Run("-a SHA512", func(t *testing.T) {
			var (
				args = []string{"hash", "-a", "sha512", "./fixtures/example.md"}
				in   = bytes.NewReader(nil)
				out  = new(bytes.Buffer)
				err  = new(bytes.Buffer)
			)

			exitCode := run(args, in, out, err)

			assertEqual(t, exitCode, 0)
			assertEqual(t, out.String(), "2c74fd17edafd80e8447b0d46741ee243b7eb74dd2149a0ab1b9246fb30382f27e853d8585719e0e67cbda0daa8f51671064615d645ae27acb15bfb1447f459b\n")
			assertEqual(t, err.String(), "")
		})
	})

	t.Run("version", func(t *testing.T) {
		t.Run("-v", func(t *testing.T) {
			var (
				args = []string{"hash", "-v"}
				in   = bytes.NewReader([]byte(""))
				out  = new(bytes.Buffer)
				err  = new(bytes.Buffer)
			)

			exitCode := run(args, in, out, err)

			assertEqual(t, exitCode, 0)
			assertEqual(t, out.String(), "hash version: test-version\n")
			assertEqual(t, err.String(), "")
		})

		t.Run("-version", func(t *testing.T) {
			var (
				args = []string{"hash", "-version"}
				in   = bytes.NewReader([]byte(""))
				out  = new(bytes.Buffer)
				err  = new(bytes.Buffer)
			)

			exitCode := run(args, in, out, err)

			assertEqual(t, exitCode, 0)
			assertEqual(t, out.String(), "hash version: test-version\n")
			assertEqual(t, err.String(), "")
		})
	})
}

func assertEqual(t *testing.T, want any, got any) {
	t.Helper()

	if got != want {
		t.Errorf("Expected %v, but got %v", want, got)
	}
}
