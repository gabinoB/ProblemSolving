package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int64) {
	var miniSum int64
	var maxSum int64
	var sumAux int64
	var kAux int64
	var c int64 = 0
	for i := 0; i < len(arr); i++ {
		for k := 0; k < len(arr); k++ {
			kAux = arr[k]
			arr[k] = 0
			for j := 0; j < len(arr); j++ {
				sumAux = sumAux + arr[j]
			}
			if c == 0 {
				miniSum = sumAux
				c++
			}
			if sumAux > maxSum {
				maxSum = sumAux
			}
			if sumAux < miniSum {
				miniSum = sumAux
			}
			arr[k] = kAux
			sumAux = 0
		}
	}
	fmt.Printf("%v %v", miniSum, maxSum)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int64

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int64(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
