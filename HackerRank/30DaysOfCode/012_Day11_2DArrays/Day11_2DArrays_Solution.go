package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func getHourglassSum(matrix [][]int32, row int, col int) int32 {
	var sum int32 = 0
	sum = sum + matrix[row-1][col-1]
	sum = sum + matrix[row-1][col]
	sum = sum + matrix[row-1][col+1]
	sum = sum + matrix[row][col]
	sum = sum + matrix[row+1][col-1]
	sum = sum + matrix[row+1][col]
	sum = sum + matrix[row+1][col+1]

	return sum
}

func getMaxHourglassSum(arr [][]int32) int32 {

	var maxHourglassSum int32 = -63

	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			currentHourglassSum := getHourglassSum(arr, j, i)
			if currentHourglassSum > maxHourglassSum {
				maxHourglassSum = currentHourglassSum
			}
		}
	}

	return maxHourglassSum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != 6 {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}
	result := getMaxHourglassSum(arr)

	fmt.Println(result)
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
