// Copyright 2020 Vasyl Naumenko. All rights reserved.

package main

import "fmt"

func countTriplets(arr []int64, r int64) int64 {
	count := int64(0)
	t := make(map[int64]int64)
	tPair := make(map[int64]int64)

	l := len(arr) - 1
	// 1 3 9 9 27 81
	for i := l; i >= 0; i-- {
		val := arr[i]
		next := r * val

		if _, ok := tPair[next]; ok {
			count += tPair[next]
		}

		if _, ok := t[next]; ok {
			if _, ok := tPair[val]; !ok {
				tPair[val] = 0
			}
			tPair[val] += t[next]
		}

		if _, ok := t[val]; !ok {
			t[val] = 0
		}
		t[val]++
	}

	return count
}

func main() {
	var arr []int64
	var r, ans int64

	// helper log
	log := func(a, b, c interface{}) {
		fmt.Printf("\n📜arr is %v \n🔺Triplets: %v should be %v\n", a, b, c)
		fmt.Printf("---\n\n")
	}

	arr = []int64{1, 2, 2, 4}
	r = int64(2)
	ans = countTriplets(arr, r)
	log(arr, ans, 2)

	arr = []int64{1, 3, 9, 9, 27, 81}
	r = int64(3)
	ans = countTriplets(arr, r)
	log(arr, ans, 6)

	arr = []int64{1, 5, 5, 25, 125}
	r = int64(5)
	ans = countTriplets(arr, r)
	log(arr, ans, 4)

	arr = make([]int64, 100)
	for i := range arr {
		arr[i] = 1
	}
	r = int64(1)
	ans = countTriplets(arr, r)
	log(arr, ans, 161700)
}
