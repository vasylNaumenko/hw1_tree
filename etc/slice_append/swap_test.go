/*
 * Copyright (c) 2021. Vasyl Naumenko
 */

package main

import "testing"

//var l = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
//var l = make([]int, 1_000)
var l = make([]int, 30)

func BenchmarkSwapSlice(b *testing.B) {
	for n := 0; n < 1000000; n++ {
		SwapSlice(l)
	}
}
func BenchmarkSwapCycle(b *testing.B) {
	for n := 0; n < 1000000; n++ {
		SwapCycle(l)
	}
}
