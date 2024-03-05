package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	exitCode := run(os.Args, os.Stdin, os.Stdout, os.Stderr)
	os.Exit(exitCode)
}

var availableAlgorthms = map[string]bool{
	"MD5":    true,
	"SHA1":   true,
	"SHA224": true,
	"SHA256": true,
	"SHA384": true,
	"SHA512": true,
}

func run(args []string, stdIn io.Reader, stdOut io.Writer, stdErr io.Writer) int {
	var (
		algo string
		help bool
	)

	f := flag.NewFlagSet("hash", flag.ExitOnError)

	f.StringVar(&algo, "a", "", "Hashing algorithm")
	f.StringVar(&algo, "algorithm", "", "Hashing algorithm")

	f.BoolVar(&help, "h", false, "Print help")
	f.BoolVar(&help, "help", false, "Print help")

	f.Usage = func() {
		usage(f.Name(), stdErr)
	}

	f.SetOutput(stdErr)
	f.Parse(args[1:])

	if help {
		f.Usage()
		return 2
	}

	if algo != "" {
		algo = strings.ToUpper(algo)
	}

	if _, exists := availableAlgorthms[algo]; exists == false {
		fmt.Fprintf(stdErr, "Unsupported hashing algorithm\n")
		return 1
	}

	switch algo {
	case "MD5":
		hasher := md5.New()
		if _, err := io.Copy(hasher, stdIn); err != nil {
			fmt.Fprint(stdErr, err)
			return 1
		}
		sum := hex.EncodeToString(hasher.Sum(nil))
		fmt.Fprintln(stdOut, sum)
	case "SHA1":
		hasher := sha1.New()
		if _, err := io.Copy(hasher, stdIn); err != nil {
			fmt.Fprint(stdErr, err)
			return 1
		}
		sum := hex.EncodeToString(hasher.Sum(nil))
		fmt.Fprintln(stdOut, sum)

	case "SHA224":
		hasher := sha256.New224()
		if _, err := io.Copy(hasher, stdIn); err != nil {
			fmt.Fprint(stdErr, err)
			return 1
		}
		sum := hex.EncodeToString(hasher.Sum(nil))
		fmt.Fprintln(stdOut, sum)

	case "SHA256":
		hasher := sha256.New()
		if _, err := io.Copy(hasher, stdIn); err != nil {
			fmt.Fprint(stdErr, err)
			return 1
		}
		sum := hex.EncodeToString(hasher.Sum(nil))
		fmt.Fprintln(stdOut, sum)

	case "SHA384":
		hasher := sha512.New384()
		if _, err := io.Copy(hasher, stdIn); err != nil {
			fmt.Fprint(stdErr, err)
			return 1
		}
		sum := hex.EncodeToString(hasher.Sum(nil))
		fmt.Fprintln(stdOut, sum)

	case "SHA512":
		hasher := sha512.New()
		if _, err := io.Copy(hasher, stdIn); err != nil {
			fmt.Fprint(stdErr, err)
			return 1
		}
		sum := hex.EncodeToString(hasher.Sum(nil))
		fmt.Fprintln(stdOut, sum)
	default:
		return 3

	}
	return 0
}

func usage(name string, stdErr io.Writer) {
	fmt.Fprintf(stdErr, `%s program usage:

hash [FLAGS] -a SHA1 -- [STDIN]
hash [FLAGS] -a SHA1 [FILE]

Flags:
-a, -algorithm one of {MD5 SHA1 SHA224 SHA256 SHA384 SHA512}
-h, -help print help
-v, -version print commit hash from which that program was built from
`, name)
}
