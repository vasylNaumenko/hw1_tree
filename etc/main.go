// Copyright 2020 Vasyl Naumenko Technologies. All rights reserved.

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf(">>>>%d \n", makeTimestampUnix())
	fmt.Printf(">>>>%d \n", makeTimestamp())
	fmt.Printf(">>>>%d \n", makeTimestampNano())

	k := makeTimestampNano()
	fmt.Printf(">>>>%d \n", k)
}

func makeTimestampUnix() int64 {
	return time.Now().Unix()

}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func makeTimestampNano() int64 {
	return time.Now().UnixNano()
}
