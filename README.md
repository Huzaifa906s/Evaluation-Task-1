# Path Normalization Utility

## Overview
This repository contains a Go program that converts a relative file path into an absolute path without using any built-in functions that perform path conversion. The solution manually processes path components to handle occurrences of `.` (current directory), `..` (parent directory), and redundant slashes.

## Task Description
The program takes:
- A **current directory** (absolute path)
- A **relative path**

It then computes the **absolute path** by resolving `.` and `..` while ensuring a normalized format.

## Features
- Manually processes path components without relying on built-in path conversion functions.
- Correctly handles:
  - `.` (current directory)
  - `..` (parent directory)
  - Multiple consecutive slashes (`//`)
- Produces a normalized absolute path without redundant elements.

## Usage
To run the program, execute the `main.go` file:

```sh
# Clone the repository
$ git clone https://github.com/yourusername/path-normalization.git
$ cd path-normalization

# Run the Go program
$ go run main.go
```

## Example
### Input:
- Current Directory: `/home/myhome/myfolder1`
- Relative Path: `../myfolder2/mydocument.txt`

### Output:
```
Absolute Path: /home/myhome/myfolder2/mydocument.txt
```

## Contributing
Contributions are welcome! Please submit a pull request or open an issue for any improvements or bug fixes.

## Author
Muhammad Huzaifa

**Email**: huzaifashamayl90@gmail.com

