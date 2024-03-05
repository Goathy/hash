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

		if exitCode != 1 {
			t.Error("Program exit with non zero code")
		}

		if want, got := "", out.String(); got != want {
			t.Errorf("Expected %v, but got %v", want, got)
		}

		if want, got := "Unsupported hashing algorithm\n", err.String(); got != want {
			t.Errorf("Expected %v, but got %v", want, got)
		}
	})

	t.Run("help", func(t *testing.T) {
		helpMsg := `hash program usage:

hash [FLAGS] -a SHA1 -- [STDIN]
hash [FLAGS] -a SHA1 [FILE]

Flags:
-a, -algorithm one of {MD5 SHA1 SHA224 SHA256 SHA384 SHA512}
-h, -help print help
-v, -version print commit hash from which that program was built from
`
		t.Run("-h", func(t *testing.T) {
			var (
				args = []string{"hash", "-h"}
				in   = bytes.NewReader([]byte(""))
				out  = new(bytes.Buffer)
				err  = new(bytes.Buffer)
				want = ""
			)

			exitCode := run(args, in, out, err)

			if exitCode != 2 {
				t.Error("Program exit with non zero code")
			}

			if got := out.String(); got != want {
				t.Errorf("Expected %v, but got %v", want, got)
			}

			if got := err.String(); got != helpMsg {
				t.Errorf("Expected %v, but got %v", helpMsg, got)
			}
		})

		t.Run("-help", func(t *testing.T) {
			var (
				args = []string{"hash", "-help"}
				in   = bytes.NewReader([]byte(""))
				out  = new(bytes.Buffer)
				err  = new(bytes.Buffer)
				want = ""
			)

			exitCode := run(args, in, out, err)

			if exitCode != 2 {
				t.Error("Program exit with non zero code")
			}

			if got := out.String(); got != want {
				t.Errorf("Expected %v, but got %v", want, got)
			}

			if got := err.String(); got != helpMsg {
				t.Errorf("Expected %v, but got %v", helpMsg, got)
			}
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

			if exitCode != 0 {
				t.Error("Program exit with non zero code")
			}

			if want, got := "b10a8db164e0754105b7a99be72e3fe5\n", out.String(); got != want {
				t.Errorf("Expected %v, but got %v", want, got)
			}

			if want, got := "", err.String(); got != want {
				t.Errorf("Expected %v, but got %v", want, got)
			}
		})

		t.Run("-algorithm SHA1", func(t *testing.T) {
			var (
				args = []string{"hash", "-algorithm", "sha1", "--"}
				in   = bytes.NewReader([]byte("Hello World"))
				out  = new(bytes.Buffer)
				err  = new(bytes.Buffer)
			)

			exitCode := run(args, in, out, err)

			if exitCode != 0 {
				t.Error("Program exit with non zero code")
			}

			if want, got := "0a4d55a8d778e5022fab701977c5d840bbc486d0\n", out.String(); got != want {
				t.Errorf("Expected %v, but got %v", want, got)
			}

			if want, got := "", err.String(); got != want {
				t.Errorf("Expected %v, but got %v", want, got)
			}
		})

		t.Run("-a SHA224", func(t *testing.T) {
			var (
				args = []string{"hash", "-a", "sha224"}
				in   = bytes.NewReader([]byte("Hello World"))
				out  = new(bytes.Buffer)
				err  = new(bytes.Buffer)
			)

			exitCode := run(args, in, out, err)

			if exitCode != 0 {
				t.Error("Program exit with non zero code")
			}

			if want, got := "c4890faffdb0105d991a461e668e276685401b02eab1ef4372795047\n", out.String(); got != want {
				t.Errorf("Expected %v, but got %v", want, got)
			}

			if want, got := "", err.String(); got != want {
				t.Errorf("Expected %v, but got %v", want, got)
			}
		})

		t.Run("-a SHA256", func(t *testing.T) {
			var (
				args = []string{"hash", "-a", "sha256", "--"}
				in   = bytes.NewReader([]byte("Hello World"))
				out  = new(bytes.Buffer)
				err  = new(bytes.Buffer)
			)

			exitCode := run(args, in, out, err)

			if exitCode != 0 {
				t.Error("Program exit with non zero code")
			}

			if want, got := "a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e\n", out.String(); got != want {
				t.Errorf("Expected %v, but got %v", want, got)
			}

			if want, got := "", err.String(); got != want {
				t.Errorf("Expected %v, but got %v", want, got)
			}
		})

		t.Run("-a SHA384", func(t *testing.T) {
			var (
				args = []string{"hash", "-a", "sha384"}
				in   = bytes.NewReader([]byte("Hello World"))
				out  = new(bytes.Buffer)
				err  = new(bytes.Buffer)
			)

			exitCode := run(args, in, out, err)

			if exitCode != 0 {
				t.Error("Program exit with non zero code")
			}

			if want,
				got := "99514329186b2f6ae4a1329e7ee6c610a729636335174ac6b740f9028396fcc803d0e93863a7c3d90f86beee782f4f3f\n",
				out.String(); got != want {
				t.Errorf("Expected %v, but got %v", want, got)
			}

			if want, got := "", err.String(); got != want {
				t.Errorf("Expected %v, but got %v", want, got)
			}
		})

		t.Run("-a SHA512", func(t *testing.T) {
			var (
				args = []string{"hash", "-a", "sha512"}
				in   = bytes.NewReader([]byte("Hello World"))
				out  = new(bytes.Buffer)
				err  = new(bytes.Buffer)
			)

			exitCode := run(args, in, out, err)

			if exitCode != 0 {
				t.Error("Program exit with non zero code")
			}

			if want,
				got := "2c74fd17edafd80e8447b0d46741ee243b7eb74dd2149a0ab1b9246fb30382f27e853d8585719e0e67cbda0daa8f51671064615d645ae27acb15bfb1447f459b\n",
				out.String(); got != want {
				t.Errorf("Expected %v, but got %v", want, got)
			}

			if want, got := "", err.String(); got != want {
				t.Errorf("Expected %v, but got %v", want, got)
			}
		})
	})
}
