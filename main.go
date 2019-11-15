package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func getFlagUsage(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	var sb strings.Builder

	for scanner.Scan() {
		sb.Write([]byte(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return sb.String()
}

func main() {
	fmt.Println("go-flag-doc")
	fmt.Println(getFlagUsage(os.Stdin))
}
