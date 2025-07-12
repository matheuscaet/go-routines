package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	OrderStatusPending   = "pending"
	OrderStatusShipped   = "shipped"
	OrderStatusDelivered = "delivered"
	OrderStatusCanceled  = "canceled"
)

type Order struct {
	ID     int
	Status string
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	// generate 50-100 orders
	orders := generateOrders(rand.Intn(50) + 50)
	fmt.Println("Orders generated:", len(orders))

	go processOrders(orders, &wg)
	go setOrdersShipped(orders, &wg)
	go setOrdersDelivered(orders, &wg)

	wg.Wait()

	fmt.Println("Orders processed:", len(orders))
}

func processOrders(orders []*Order, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, order := range orders {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Println("Order", order.ID, "processed")
	}
}

func setOrdersShipped(orders []*Order, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, order := range orders {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Println("Order", order.ID, "shipped")
		order.changeStatus(OrderStatusShipped)
	}
}

func setOrdersDelivered(orders []*Order, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, order := range orders {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Println("Order", order.ID, "delivered")
		order.changeStatus(OrderStatusDelivered)
	}
}

func generateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := range orders {
		orders[i] = &Order{
			ID:     i + 1,
			Status: OrderStatusPending,
		}
	}
	return orders
}

func (o *Order) changeStatus(status string) {
	o.Status = status
}
