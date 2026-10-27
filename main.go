package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Status string
}

var ordersMu sync.RWMutex

func main() {
	orders := generateOrders(20)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		processOrders(orders)
	}()

	go func() {
		defer wg.Done()
		updateOrderStatus(orders)
	}()

	go func() {
		defer wg.Done()
		reportOrderStatus(orders)
	}()

	wg.Wait()
	fmt.Println("All operations completed.")
}

func generateOrders(count int) []*Order {
	orders := make([]*Order, count)

	for i := range count {
		orders[i] = &Order{
			ID:     i + 1,
			Status: "Pending",
		}
	}

	return orders
}

func processOrders(orders []*Order) {
	for _, order := range orders {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Printf("Processing order %d\n", order.ID)
	}
}

func updateOrderStatus(orders []*Order) {
	statuses := []string{"Processing", "Shipped", "Delivered"}

	for _, order := range orders {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		status := statuses[rand.Intn(len(statuses))]

		ordersMu.Lock()
		order.Status = status
		ordersMu.Unlock()

		fmt.Printf("Updated order %d: %s\n", order.ID, status)
	}
}

func reportOrderStatus(orders []*Order) {
	counts := make(map[string]int)

	for _, order := range orders {
		ordersMu.RLock()
		status := order.Status
		ordersMu.RUnlock()

		counts[status]++
		fmt.Printf("Order %d: %s\n", order.ID, status)
	}

	fmt.Println("\n--- Order Status Summary ---")
	for _, status := range []string{
		"Pending", "Processing", "Shipped", "Delivered",
	} {
		fmt.Printf("%-12s %d\n", status, counts[status])
	}
}
