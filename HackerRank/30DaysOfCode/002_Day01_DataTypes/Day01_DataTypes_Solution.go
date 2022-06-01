package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.

	// Read and save an integer, double, and String to your variables.

	// Print the sum of both integer variables on a new line.

	// Print the sum of the double variables on a new line.

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.

	var j uint64
	var f float64
	var res float64

	fmt.Scanln(&j)
	fmt.Scanln(&f)

	res = d + f

	fmt.Println(i + j)
	fmt.Printf("%.1f\n", res)

	for scanner.Scan() {
		fmt.Printf(s + scanner.Text() + "\n")
	}
}
