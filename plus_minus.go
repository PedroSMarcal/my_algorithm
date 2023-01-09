package main

import "fmt"

// Enunciado

/*
Given an array of integers, calculate the ratios of its elements that are positive, negative, and zero. Print the decimal value of each fraction on a new line with 6 places after the decimal.

Note: This challenge introduces precision problems. The test cases are scaled to six decimal places, though answers with absolute error of up 10^-4 to are acceptable.
*/

func plusMinus(arr []int32) {
	tam := len(arr)
	var response [3]int32

	for _, value := range arr {
		if value > 0 {
			response[0]++
			continue
		} else if value == 0 {
			response[2]++
			continue
		}
		response[1]++
		continue
	}

	for _, value := range response {
		res := float32(value) / float32(tam)
		fmt.Println(res)
	}
}
