package main

import (
	"fmt"
	"strconv"
)

// get all prime numbers until the got number

func main() {
	var N int
	fmt.Scan(&N)
	ans := ""
	for i := 1; i <= N; i++ {
		toAdd := true
		for j := 2; j <= N; j++ {
			if i != j && i%j == 0 {
				toAdd = false
			}
		}
		if toAdd {
			ans += strconv.Itoa(i) + "\n"
		}
	}
	fmt.Printf(ans)
}
