package main

import (
	"github.com/lfrei/kafka-store/controller"
	"github.com/lfrei/kafka-store/messaging"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go messaging.Start(&wg, "product")
	go controller.Start(&wg, "product")

	wg.Wait()
}
