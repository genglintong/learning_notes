package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string] int)
	files := os.Args[1:]
	if len(files) <= 0 {
		return
	}

	for _, arg := range files {
		f , err := os.Open(arg)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}

		countLines(f, counts)
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int)  {
	input := bufio.NewScanner(f)

	for input.Scan() {
		s := input.Text()
		if counts[s] > 0 {
			fmt.Printf("this file has dup data! fileName: %s\n", f.Name())
		}
		counts[s]++
	}
}
