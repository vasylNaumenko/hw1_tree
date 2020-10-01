package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// Simple structure which tags chef and order number
type chefOrder struct {
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

// Function where the waiter takes and brings your order
func processOrder(takes string, order int, orderChan chan chefOrder) {
	chef, waiter := selectWaiterAndChef()
	if takes == "takes" {
		fmt.Printf("Waiter %s takes order %d to chef %s \n", waiter, order, chef)
		orderChan <- chefOrder{chef, order}
	} else {
		fmt.Printf("Waiter %s brings order %d from chef %s \n", waiter, order, chef)
	}
}

// Function where the Chef cooks your meal
func chefCooks(orderChan chan chefOrder) {
	gotOrder := <-orderChan
	fmt.Printf("Chef %s cooks order %d \n", gotOrder.chef, gotOrder.order)
}

// Main function
func main() {
	runtime.GOMAXPROCS(1)
	totalOrders := 5

	orderChan := make(chan chefOrder)

	for i := 0; i < totalOrders; i++ {
		go processOrder("takes", i, orderChan)
		go chefCooks(orderChan)
		go processOrder("brings", i, orderChan)
		time.Sleep(1)
	}
	<-time.After(time.Second * 5)
}
