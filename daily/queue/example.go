package main

import "fmt"

func main() {
	queue := make([] int ,0,5)

	queue = append(queue, 1)

	queue = queue[1:]
	fmt.Printf("len(queue): %v, cap(queue): %v, queue: %v\n", len(queue), cap(queue), queue)
}
