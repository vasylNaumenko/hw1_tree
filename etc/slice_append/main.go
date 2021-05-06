/*
 * Copyright (c) 2021. Vasyl Naumenko
 */

package main

import "fmt"

func main() {
	fmt.Println("hello")
	// SwapSlice([]int{1, 2, 3,4, 5})
	SwapCycle([]int{1, 2, 3, 4, 5})
}

func SwapSlice(a []int) {
	for i := 0; i < len(a); i++ {
		a = append(a[1:], a[0])
		//fmt.Printf("%v cap=%v\n", a, cap(a))
	}
}

func SwapCycle(a []int) {
	for i := 0; i < len(a); i++ {
		f := a[0]
		for j := 1; j < len(a); j++ {
			a[j-1] = a[j]
		}
		a[len(a)-1] = f
		//fmt.Printf("%v cap=%v\n", a, cap(a))
	}
}
