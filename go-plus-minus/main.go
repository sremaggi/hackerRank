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
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	//Counters
	var minus, plus, zero int32

	lenArr := int32(len(arr))
	for _, v := range arr {
		switch {
		case v > 0:
			plus += 1
		case v < 0:
			minus += 1
		case v == 0:
			zero += 1

		}

	}

	plusResult := float32(plus) / float32(lenArr)
	minusResult := float32(minus) / float32(lenArr)
	zeroResult := float32(zero) / float32(lenArr)

	fmt.Printf("%.6f \n", plusResult)
	fmt.Printf("%.6f \n", minusResult)
	fmt.Printf("%.6f \n", zeroResult)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)
	if n < 1 || n > 100 {
		panic("Invalid length of array")
	}

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {

		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		if arrItem < -100 || arrItem > 100 {
			panic("Invalid size of array element")
		}
		arr = append(arr, arrItem)
	}

	plusMinus(arr)

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
