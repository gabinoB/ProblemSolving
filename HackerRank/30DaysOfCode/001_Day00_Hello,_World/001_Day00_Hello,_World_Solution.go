package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')

	fmt.Println("Hello, World.")
	fmt.Printf(line)
}
