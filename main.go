package main

import (
	"fmt"
	"strings"
)

func normalizePath(currentDir, relativePath string) string {

	var currentParts, relativeParts []string
	for _, part := range strings.Split(currentDir, "/") {
		if part != "" {
			currentParts = append(currentParts, part)
		}
	}
	for _, part := range strings.Split(relativePath, "/") {
		if part != "" {
			relativeParts = append(relativeParts, part)
		}
	}

	var stack []string

	stack = append(stack, currentParts...)

	for _, part := range relativeParts {
		switch part {
		case ".":

			continue
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, part)
		}
	}

	return "/" + strings.Join(stack, "/")
}

func main() {
	currentDir := "/home/myhome/myfolder1"
	relativePath := "./mydocument.txt"

	absolutePath := normalizePath(currentDir, relativePath)
	fmt.Println("Absolute Path:", absolutePath)
}
