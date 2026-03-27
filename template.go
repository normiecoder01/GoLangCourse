package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID int
}

func main() {

	orders := make(chan Order)          // input
	validated := make(chan Order, 5)    // buffered
	processed := make(chan Order, 5)    // buffered

	var wg sync.WaitGroup

	// -------- PRODUCER --------
	go func() {
		for i := 1; i <= 10; i++ {
			order := Order{ID: i}
			fmt.Println("Created Order:", order.ID)
			orders <- order
		}
		close(orders) // no more incoming orders
	}()

	// -------- VALIDATION WORKERS --------
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			for order := range orders {
				fmt.Println("Validator", workerID, "processing order", order.ID)
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))

				validated <- order
			}
		}(i)
	}

	// Close validated channel AFTER all validators finish
	go func() {
		wg.Wait()
		close(validated)
	}()

	// -------- PROCESSING WORKERS --------
	var wg2 sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg2.Add(1)
		go func(workerID int) {
			defer wg2.Done()

			for order := range validated {
				fmt.Println("Processor", workerID, "processing payment for order", order.ID)
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))

				processed <- order
			}
		}(i)
	}

	// Close processed channel AFTER processing
	go func() {
		wg2.Wait()
		close(processed)
	}()

	// -------- FINAL CONSUMER --------
	for order := range processed {
		fmt.Println("Shipping order", order.ID)
	}

	fmt.Println("All orders completed")
}