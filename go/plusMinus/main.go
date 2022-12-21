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
Given an array of ints, positive, negative or 0
Return ratios of postive/length, negative/length, 0-value/length
Input: [10, 30, 0, -23, -45]
Return: 0.400000, 0.400000, 0.200000
*/
func plusMinus(arr []int32) (float64, float64, float64) {
	length := float64(len(arr))
	var positives, negatives, zeros float64 = 0, 0, 0
	for _, v := range arr {
		if v > 0 {
			positives++
		} else if v < 0 {
			negatives++
		} else {
			zeros++
		}
	}
	return positives / length, negatives / length, zeros / length
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	possRatio, negsRatio, zerosRatios := plusMinus(arr)
	fmt.Println(fmt.Sprintf("%.6f\n%.6f\n%.6f\n", possRatio, negsRatio, zerosRatios))
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
