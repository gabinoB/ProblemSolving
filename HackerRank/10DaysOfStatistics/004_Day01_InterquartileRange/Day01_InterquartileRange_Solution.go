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
 * Complete the 'interQuartile' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY values
 *  2. INTEGER_ARRAY freqs
 */

func createDataSet(values []int32, freqs []int32) []int32 {

	var S []int32

	for val := range values {
		for i := 0; i < int(freqs[i]); i++ {
			S = append(S, int32(val))
		}
	}

	return S
}

func interQuartile(values []int32, freqs []int32) {
	// Print your answer to 1 decimal place within this function

	/*
		Crear un array S en el que haya tantas veces el primer valor de values como indique el primer valor de freqs, para esto tengo que hacer que, con un bucle for que
		recorre el arr values se repita otro bucle for la cantidad de veces que indica el valor de freqs[i]

		Quizas con un bucle for "i j" se pueda hacer.
	*/

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
