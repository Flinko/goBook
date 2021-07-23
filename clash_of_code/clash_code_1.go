package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	//fmt.Printf("number: %d", N)
	ans := 0
	for i := 0; i < N; i++ {
		var A, B int
		fmt.Scan(&A, &B)
		fmt.Printf("A: %d, B: %d\n", A, B)
		ans = (A * A) + B
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Printf("answer: %d\n", ans) // Write answer to stdout
}
