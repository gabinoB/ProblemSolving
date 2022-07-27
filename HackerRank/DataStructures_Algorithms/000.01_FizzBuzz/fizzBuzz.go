package main

import "fmt"

func main() {
	var num int64
	fmt.Print("Insert number:")
	fmt.Scan(&num)

	for i := 0; int64(i) <= num; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Fizz")
		} else if i%3 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
