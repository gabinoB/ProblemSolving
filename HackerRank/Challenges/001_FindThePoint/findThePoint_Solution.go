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
 * Complete the 'findPoint' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER px
 *  2. INTEGER py
 *  3. INTEGER qx
 *  4. INTEGER qy
 */

func findPoint(px int32, py int32, qx int32, qy int32) []int32 {
    // Write your code here
    
    //r is the result, the variable to return
    r := make([]int32,2)
    
    //if the point are in the same position. r gona be the exact same position to p and q
    if px == py && px == qx && px == qy {
        r[0] = px
        r[1] = py
    } else if px <= qx && py <= qy { 
        r[0] = qx + (qx - px)
        r[1] = qy + (qy - py)
    } else if px <= qx && py >= qy {
        r[0] = qx + (qx - px)
        r[1] = qy - (py - qy)
    } else if px >= qx && py >= qy {
        r[0] = qx - (px - qx)
        r[1] = qy - (py - qy)
    } else if px >= qx && py <= qy {
        r[0] = qx - (px - qx)
        r[1] = qy + (qy - py)
    }
    return r
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    for nItr := 0; nItr < int(n); nItr++ {
        firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        pxTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
        checkError(err)
        px := int32(pxTemp)

        pyTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
        checkError(err)
        py := int32(pyTemp)

        qxTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
        checkError(err)
        qx := int32(qxTemp)

        qyTemp, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
        checkError(err)
        qy := int32(qyTemp)

        result := findPoint(px, py, qx, qy)

        for i, resultItem := range result {
            fmt.Fprintf(writer, "%d", resultItem)

            if i != len(result) - 1 {
                fmt.Fprintf(writer, " ")
            }
        }

        fmt.Fprintf(writer, "\n")
    }

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
