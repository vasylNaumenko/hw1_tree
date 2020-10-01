package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// Simple structure which tags chef and order numbe
type chefOrder struct {
	chef  string
	order int
}

// Same as above, duplex channel
type chefCooked struct {
	chef  string
	order int
}

// Function to select a Chef and Waiter for you
func selectWaiterAndChef() (string, string) {
	allChefs := []string{"Jack", "Bob", "Mark"}
	allWaiters := []string{"A", "B", "C"}

	randomC := rand.Intn(len(allChefs))
	randomW := rand.Intn(len(allWaiters))
	cpick := allChefs[randomC]
	wpick := allWaiters[randomW]

	return cpick, wpick
}

// Function where the waiter takes order, chef cooks it and waiter brings your order
func processOrderandCook(order int, orderChan chan chefOrder, cookChan chan chefCooked) {
	chef, waiter := selectWaiterAndChef()
	select {
	case cookedOrder := <-cookChan:
		fmt.Printf("Waiter %s brings order %d from chef %s \n", waiter, cookedOrder.order, cookedOrder.chef)
	case gotOrder := <-orderChan:
		fmt.Printf("Chef %s cooks order %d \n", gotOrder.chef, gotOrder.order)
		cookChan <- chefCooked{gotOrder.chef, gotOrder.order}
	default:
		fmt.Printf("Waiter %s takes order %d to chef %s \n", waiter, order, chef)
		orderChan <- chefOrder{chef, order}
	}
}

// Main function
func main() {
	runtime.GOMAXPROCS(1)
	totalOrders := 5

	orderChan := make(chan chefOrder)
	cookChan := make(chan chefCooked)

	for i := 0; i < totalOrders; i++ {
		go processOrderandCook(i, orderChan, cookChan)
		// lol
		go processOrderandCook(i, orderChan, cookChan)
		go processOrderandCook(i, orderChan, cookChan)
	}
	<-time.After(time.Second * 5)
}
