package main

import (
	"fmt"
)

// Enunciado

/*
Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly four of the five integers. Then print the respective minimum and maximum values as a single line of two space-separated long integers.
*/

// func sortArray(arr []int32) {
// 	var index int
// 	var safe int32
// 	arrLength := len(arr) - 1
// 	for i := arrLength; i <= arrLength; i++ {
// 		for j := arrLength; j <= arrLength; j++ {

// 		}
// 	}
// }

func minMax(arr []int32, value *int32) {

	for i := 0; i <= 3; i++ {
		*value += arr[i]
	}
}

func getMin(arr []int32, min *int32) {

	minMax(arr, min)
}

func getMax(arr []int32, max *int32) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	minMax(arr, max)
}

func miniMaxSum(arr []int32) {

	var min int32
	var max int32

	getMin(arr, &min)
	getMax(arr, &max)

	fmt.Printf("%v %v", min, max)
}
