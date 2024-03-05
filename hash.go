package main

import (
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

func run(args []string, in io.Reader, out io.Writer, stdErr io.Writer) int {
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
