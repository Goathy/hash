# Hash Utility

A simple, yet effective, command-line utility written in Go that facilitates hashing of input data using various cryptographic algorithms. This utility allows users to hash data from standard input (STDIN) or from a file, using their chosen hash function. It supports a variety of cryptographic hash functions including MD5, SHA1, SHA224, SHA256, SHA384, and SHA512.

## Features

- **Multiple Hashing Algorithms**: Supports MD5, SHA1, SHA224, SHA256, SHA384, and SHA512 hash algorithms for versatile cryptographic needs.
- **CLI Convenience**: Offers a straightforward Command Line Interface for easy operation.
- **Input Flexibility**: Accepts input either from STDIN or a specified file, enabling hashing of diverse data sources.
- **Error Handling**: Provides meaningful error messages for incorrect usage or operational failures.

## Installation

To run this utility, ensure that you have Go installed on your system. This program doesn't require any external packages beyond the Go standard library.

1. Clone the repository or download the source code.
2. Navigate to the folder containing the code.

## Usage

To use the utility, invoke it from the command line. Below are the available flags and an example command:

```
Flags:
-a, -algorithm   Specify hashing algorithm {MD5, SHA1, SHA224, SHA256, SHA384, SHA512}
-h, -help        Print the help message
```

### Examples

- Hashing input from STDIN using SHA256:

  ```sh
  echo -n "Hello, World!" | go run . -a SHA256
  ```

- Hashing the contents of a file using SHA1:

  ```sh
  go run . -a SHA1 path/to/your/file.txt
  ```

**Note**: The utility reads from STDIN if no file path is specified or if the file path is "--".

## Contributing

Contributions to enhance the utility, add new features, or fix bugs are always welcome. Please feel free to fork the repository and submit a pull request.

## License

This project is open-sourced under the MIT License. See the LICENSE file for more details.
