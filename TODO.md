Expected program usage:

hash [FLAGS] -i [FILE]
hash [FLAGS] -- [STDIN]

Flags:
-a, -algorithm one of {MD5 SHA1 SHA224 SHA256 SHA384 SHA512}
-i, -input path to file to hash
-h, -help print help
-v, -version print commit hash from which that program was built from

Resources:
https://pkg.go.dev/crypto@go1.22.0
https://pkg.go.dev/flag@go1.22.0
