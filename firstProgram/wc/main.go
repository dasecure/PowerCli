package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)
	// scanner.Split(bufio.ScanWords)

	wc := 0
	if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
	} else if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	for scanner.Scan() {
		wc++
	}

	return wc
}

func main() {
	lines := flag.Bool("l", false, "count lines")
	bytes := flag.Bool("c", false, "count bytes")
	// words := flag.Bool("w", false, "count words")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *bytes))
}
