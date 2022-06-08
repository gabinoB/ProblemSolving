package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'quartiles' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func evenLowQuartile(arr []int32) int32 {
	if (len(arr)/2)%2 == 0 {
		return ((arr[((len(arr)/2)/2)-1]) + (arr[((len(arr)/2)/2)])/2)
	} else {
		return (arr[((len(arr) / 2) / 2)])
	}
}

func evenMidQuartile(arr []int32) int32 {
	return ((arr[len(arr)/2] + arr[(len(arr)/2)+1]) / 2)
}

func reverseArray(arr []int32) []int32 {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func evenHighQuartile(arr []int32) int32 {

	inverseArr := reverseArray(arr)

	numToReturn := evenLowQuartile(inverseArr)

	return numToReturn
}

func oddLowQuartile(arr []int32) int32 {
	if (len(arr)/2)%2 == 0 {
		return ((arr[(((len(arr))/2)/2)] + arr[(((len(arr))/2)/2)-1]) / 2)
	} else {
		return (arr[((len(arr))/2)/2])
	}
}

func oddMidQuartile(arr []int32) int32 {
	return arr[(len(arr) / 2)]
}

func oddHighQuartile(arr []int32) int32 {
	inversedArr := reverseArray(arr)

	return oddLowQuartile(inversedArr)
}

func quartiles(arr []int32) [3]int32 {
	// Write your code here

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	var lowQuartile int32
	var midQuartile int32
	var highQuartile int32
	var solution_ArrToReturn [3]int32

	if len(arr) == 10 {
		lowQuartile = evenLowQuartile(arr)
		midQuartile = evenMidQuartile(arr)
		highQuartile = evenHighQuartile(arr)
	} else {
		lowQuartile = oddLowQuartile(arr)
		midQuartile = oddMidQuartile(arr)
		highQuartile = oddHighQuartile(arr)
	}

	solution_ArrToReturn[0] = lowQuartile
	solution_ArrToReturn[1] = midQuartile
	solution_ArrToReturn[2] = highQuartile

	return solution_ArrToReturn
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	dataTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var data []int32

	for i := 0; i < int(n); i++ {
		dataItemTemp, err := strconv.ParseInt(dataTemp[i], 10, 64)
		checkError(err)
		dataItem := int32(dataItemTemp)
		data = append(data, dataItem)
	}

	res := quartiles(data)

	for i, resItem := range res {
		fmt.Fprintf(writer, "%d", resItem)

		if i != len(res)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
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
