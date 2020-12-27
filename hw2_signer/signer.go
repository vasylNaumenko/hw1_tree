package main

import (
	"fmt"
	"sync"
)

// сюда писать код

func CombineResults(in, out chan interface{}) {
	// for data := range in {
	// 	fmt.Printf("CombineResults data %#v", data)
	// 	out<- 3
	// }
	// in<-out
}

func MultiHash(in, out chan interface{}) {
	// for data := range in {
	// 	fmt.Printf("CombineResults data %#v", data)
	// 	out<- 2
	// }

}

func SingleHash(in, out chan interface{}) {
	var st string
	l := len(out)
	fmt.Printf("SingleHash len %#v\n", l)
	for i := 0; i < l; i++ {
		data := <-out
		fmt.Printf("SingleHash data %#v\n", data)
		st += "s_"
	}
	// data := <- out
	// fmt.Printf("SingleHash data %#v\n", data)
}

func ExecutePipeline(jobs ...job) {
	in := make(chan interface{}, MaxInputDataLen)
	out := make(chan interface{}, MaxInputDataLen)

	for i, worker := range jobs {
		fmt.Printf("Waiting...%v\n", i)

		worker(in, out)
		go func() {

			l := len(out)
			for i := 0; i < l; i++ {
				o := <-out
				fmt.Printf("= %v %v %v \n", o, i, l)
				in <- o
			}

			fmt.Printf("closed end\n")
		}()
	}
	fmt.Printf("Waiting...\n")

	fmt.Printf("DONE WAIT\n")
}

func process(worker job, in chan interface{}, wg *sync.WaitGroup) chan interface{} {
	out := make(chan interface{}, MaxInputDataLen)
	worker(in, out)
	go func() {
		defer wg.Done()
		l := len(out)
		for i := 0; i < l; i++ {
			o := <-out
			fmt.Printf("= %v %v %v \n", o, i, l)
			in <- o
		}
		//close(out)
		fmt.Printf("closed end\n")
	}()

	return out
}

//
// func ExecutePipeline(jobs ...job) {
// 	in := make(chan interface{}, MaxInputDataLen)
//
// 	var wg sync.WaitGroup
// 	chans := make(map[int]chan interface{})
// 	chans[0] = in
// 	for i, worker := range jobs {
//
// 		wg.Add(1)
// 		fmt.Printf("%v\n", i)
//
// 		chans[i+1] = process(worker,chans[i], &wg)
// 		out := make(chan interface{}, MaxInputDataLen)
//
// 	}
// 	fmt.Printf("Waiting...\n")
// 	wg.Wait()
// 	fmt.Printf("DONE WAIT\n")
// }
//
// func process(worker job, in chan interface{}, wg *sync.WaitGroup) chan interface{} {
// 	out := make(chan interface{}, MaxInputDataLen)
// 	worker(in, out)
// 	go func() {
// 		defer wg.Done()
// 		l := len(out)
// 		for i := 0; i < l; i++ {
// 			o := <-out
// 			fmt.Printf("= %v %v %v \n", o, i, l)
// 			in <- o
// 		}
// 		//close(out)
// 		fmt.Printf("closed end\n")
// 	}()
//
// 	return out
// }
