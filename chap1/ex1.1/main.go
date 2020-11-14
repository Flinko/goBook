package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var s string
	for idx, arg := range os.Args {
		s += strconv.Itoa(idx) + " " + arg + "\n"
	}
	fmt.Println(s)
	start := time.Now()
	ex13(os.Args)
	end := time.Now().Sub(start)
	fmt.Printf("regular took: %v\n", end.Nanoseconds())
	start = time.Now()
	ex13Join(os.Args)
	end = time.Now().Sub(start)
	fmt.Printf("with join took: %v\n", end.Nanoseconds())
}

func ex13(args []string) {
	var s string
	for idx, arg := range args {
		s += strconv.Itoa(idx) + " " + arg + "\n"
	}
	fmt.Println(s)
}

//if we a lot of values strings.Join more efficient them +=
func ex13Join(args []string) {
	var s string
	for idx, arg := range args {
		s = strings.Join([]string{s, strconv.Itoa(idx), arg}, " ")
	}
	fmt.Println(s)
}
