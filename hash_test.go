package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestHash(t *testing.T) {
	t.Run("unsupported hashing algorithm", func(t *testing.T) {
		var (
			args = []string{"hash", "-a", "md4"}
			in   = strings.NewReader("")
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
				in   = strings.NewReader("")
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
				in   = strings.NewReader("")
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
}
