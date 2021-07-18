package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "os.Stdin", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts)
			err = f.Close()
			if err != nil {
				fmt.Printf("failed to close the file: %s, err: %s", arg, err.Error())
				return
			}
		}
	}

	for line, filenames := range counts {
		for fn, ct := range filenames {
			if ct > 1 {
				fmt.Println(line, "is duplicated in ", fn, ct, "times")
			}
		}
	}
}

func countLines(f *os.File, filename string, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][filename]++
	}
}
