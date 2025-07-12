# Goroutines Example: Order Processing System

This repository demonstrates the use of Go's goroutines and synchronization primitives to simulate an order processing system. The example showcases how to handle concurrent tasks such as processing, shipping, and delivering orders using goroutines and a WaitGroup.

## Overview

The main components of the example are:

- **Order struct**: Represents an order with an ID and a status.
- **Goroutines**: Three concurrent goroutines simulate processing, shipping, and delivering orders.
- **WaitGroup**: Ensures the main function waits for all goroutines to finish before exiting.

## How It Works

1. **Order Generation**:  
   The program generates a list of orders, each starting with a "pending" status.

2. **Concurrent Processing**:  
   - `processOrders`: Simulates processing each order.
   - `setOrdersShipped`: Simulates shipping each order and updates its status.
   - `setOrdersDelivered`: Simulates delivering each order and updates its status.

3. **Synchronization**:  
   A `sync.WaitGroup` is used to wait for all three goroutines to complete their work.

## Running the Example

To run the example, use the provided Makefile:
