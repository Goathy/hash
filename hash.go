package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	exitCode := run(os.Args, os.Stdin, os.Stdout, os.Stderr)
	os.Exit(exitCode)
}

var availableAlgorthms = map[string]bool{
	"md5": true,
}

func run(args []string, in io.Reader, out io.Writer, err io.Writer) int {
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
		fmt.Fprintf(err, `%s program usage:

hash [FLAGS] -a SHA1 -i [FILE]
hash [FLAGS] -a SHA1 -- [STDIN]

Flags:
-a, -algorithm one of {MD5 SHA1 SHA224 SHA256 SHA384 SHA512}
-i, -input path to file to hash
-h, -help print help
-v, -version print commit hash from which that program was built from
`, f.Name())
	}

	f.SetOutput(err)
	f.Parse(args[1:])

	if help {
		f.Usage()
		return 2
	}

	if _, exists := availableAlgorthms[algo]; exists == false {
		fmt.Fprintf(err, "Unsupported hashing algorithm\n")
		return 1
	}

	return 0
}
