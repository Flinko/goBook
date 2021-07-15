package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	result := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		return
	} else {
		for _, arg := range files {
			counts := make(map[string]int)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			err = f.Close()
			if err != nil {
				return
			}
			for line, count := range counts {
				if count > 1 {
					result[arg] = result[arg] + line + " "
				}
			}
		}
	}
	for file, lines := range result {
		fmt.Printf("%s\t%s\n", file, lines)
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
