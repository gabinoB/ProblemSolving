package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

func interQuartile(values []int32, freqs []int32) {

	S := createDataSet(values, freqs)

	fmt.Println(S)

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
