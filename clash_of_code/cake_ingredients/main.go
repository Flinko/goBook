package main

import (
	"fmt"
	"strconv"
)
import "os"
import "bufio"
import "strings"

/**
 * "CaKe;102,108,111,117,114+101,103,103,115+109,105,108,107" => "Cake ingredients are flour, eggs, milk"
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	data := scanner.Text()
	_ = data // to avoid unused error
	res := ""
	splitted := strings.Split(data, ";")
	res += strings.ToLower(splitted[0]) + " ingredients are "
	for _, str := range strings.Split(splitted[1], ",") {
		if strings.Contains(str, "+") {
			strs := strings.Split(str, "+")
			num1, _ := strconv.Atoi(strs[0])
			r1 := rune(num1)
			num2, _ := strconv.Atoi(strs[1])
			r2 := rune(num2)
			res += string(r1) + "," + string(r2)
			continue
		}
		num, _ := strconv.Atoi(str)
		r := rune(num)
		res += string(r)
	}
	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(res) // NEVER GIVEUP :)
}

//CakE;102,108,111,117,114+101,103,103,115+109,105,108,107
