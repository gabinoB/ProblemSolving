package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func intToBinary(n int32) []int32 {

	var bin []int32
	for n >= 1 {
		bin = append(bin, n%2)
		n = n / 2
	}
	for i := 0; i < len(bin)/2; i++ {
		j := len(bin) - i - 1
		bin[i], bin[j] = bin[j], bin[i]
	}
	return bin
}

func maxConsecutives(bin []int32) int32 {
	var n int32
	var m int32 = 0
	for i := 0; i < len(bin); i++ {
		if bin[i] == 1 {
			n++
		} else if bin[i] == 0 {
			if n > m {
				m = n
			}
			n = 0
		}
	}
	if n > m {
		return n
	} else {
		return m
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)
	bin := intToBinary(n)
	res := maxConsecutives(bin)
	fmt.Println(res)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
