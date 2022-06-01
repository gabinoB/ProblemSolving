package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	data := make(map[string]string)
	for i := 0; i < n; i++ {
		scanner.Scan()
		data[strings.Split(scanner.Text(), " ")[0]] = strings.Split(scanner.Text(), " ")[1]
	}
	for scanner.Scan() {
		if data[scanner.Text()] == "" {
			fmt.Println("Not found")
		} else {
			fmt.Printf("%s=%s\n", scanner.Text(), data[scanner.Text()])
		}
	}
}
