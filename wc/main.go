package main

import (
	// read text
	"bufio"
	// print text
	"fmt"
	// io.Reader interface
	"io"
	// OS resources
	"os"
)

func count(r io.Reader) int {
	// A scanner is used to read text from a Reader (such as file)
	scanner := bufio.NewScanner(r)

	// Define the scanner split type to words (default is split by lines)
	scanner.Split(bufio.ScanWords)

	// Defining a counter
	wc := 0

	// For every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	return wc
}

func main() {
	fmt.Println(count(os.Stdin))
}
