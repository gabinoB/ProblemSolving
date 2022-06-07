package main

import (
	"fmt"
	"sort"
)

func calculateMean(arr []float64) float64 {

	var sum float64
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	return sum / float64(len(arr))
}

func calculateMedian(arr []float64) float64 {

	sort.Float64s(arr)
	m := len(arr) / 2
	res := (arr[m] + arr[m-1]) / 2
	return res
}

func findMode(arr []float64) float64 {

	type Element struct {
		id      int
		appears int
	}

	var arrayOfElements []Element

	var element Element
	element.id = int(arr[0])
	element.appears = 1

	for i := 1; i < len(arr); i++ {
		if int(arr[i]) == element.id {
			element.appears = element.appears + 1
		} else {
			arrayOfElements = append(arrayOfElements, element)
			element.id = int(arr[i])
			element.appears = 1
		}
	}

	result := arrayOfElements[0].id
	countOfResult := arrayOfElements[0].appears
	for i := 1; i < len(arrayOfElements); i++ {
		if arrayOfElements[i].appears > countOfResult {
			result = arrayOfElements[i].id
			countOfResult = arrayOfElements[i].appears
		}
	}

	return float64(result)
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT

	var n float64
	fmt.Scan(&n)
	arr := make([]float64, 0, int(n))
	for i := 0; float64(i) < n; i++ {
		var num float64
		fmt.Scan(&num)
		arr = append(arr, num)
	}

	mean := calculateMean(arr)
	fmt.Printf("%.1f\n", mean)

	median := calculateMedian(arr)
	fmt.Printf("%.1f\n", median)

	mode := findMode(arr)
	fmt.Println(mode)
}
