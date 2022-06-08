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

func createDataSet(values []int32, freqs []int32) []int32 {

	var S []int32

	for i := 0; i < len(values); i++ {
		v := values[i]
		for j := 0; j < int(freqs[i]); j++ {
			S = append(S, v)
		}
	}

	return S
}

func getLowQuartileForEvenLenArr(arr []int32) int32 {
	if (len(arr)/2)%2 == 0 {
		return (((arr[((len(arr)/2)/2)-1]) + (arr[((len(arr) / 2) / 2)])) / 2)
	} else {
		return (arr[((len(arr) / 2) / 2)])
	}
}

func reverseArray(arr []int32) []int32 {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func getLowQuartileForOddLenArr(arr []int32) int32 {
	if (len(arr)/2)%2 == 0 {
		return ((arr[(((len(arr))/2)/2)] + arr[(((len(arr))/2)/2)-1]) / 2)
	} else {
		return (arr[((len(arr))/2)/2])
	}
}

func interQuartile(values []int32, freqs []int32) {

	S := createDataSet(values, freqs)

	sort.Slice(S, func(i, j int) bool {
		return S[i] < S[j]
	})

	var lowQuartile int32
	var highQuartile int32

	if len(S)%2 == 0 {
		lowQuartile = getLowQuartileForEvenLenArr(S)
		reversedS := reverseArray(S)
		highQuartile = getLowQuartileForEvenLenArr(reversedS)
	} else {
		lowQuartile = getLowQuartileForOddLenArr(S)
		reversedS := reverseArray(S)
		highQuartile = getLowQuartileForOddLenArr(reversedS)
	}

	res := highQuartile - lowQuartile

	fmt.Printf("%.1f", float64(res))

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	valTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var val []int32

	for i := 0; i < int(n); i++ {
		valItemTemp, err := strconv.ParseInt(valTemp[i], 10, 64)
		checkError(err)
		valItem := int32(valItemTemp)
		val = append(val, valItem)
	}

	freqTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var freq []int32

	for i := 0; i < int(n); i++ {
		freqItemTemp, err := strconv.ParseInt(freqTemp[i], 10, 64)
		checkError(err)
		freqItem := int32(freqItemTemp)
		freq = append(freq, freqItem)
	}

	interQuartile(val, freq)
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
