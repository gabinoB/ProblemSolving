package main

import (
	"fmt"
	"strings"
)

func arrToString(strArray []string) string {
	return strings.Join(strArray, "")
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT

	var T int
	var S string
	var even, odd []string
	var even2, odd2 string

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&S)

		for j := 0; j < len(S); j = j + 2 {
			even = append(even, string(S[j]))
		}

		for k := 1; k < len(S); k = k + 2 {
			odd = append(odd, string(S[k]))
		}

		even2 = arrToString(even)
		odd2 = arrToString(odd)

		fmt.Printf("%s %s\n", even2, odd2)
		S = ""
		even = nil
		even2 = ""
		odd = nil
		odd2 = ""
	}

}
