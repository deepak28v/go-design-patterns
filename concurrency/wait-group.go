package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func waitGroupDemo() {
	var wait sync.WaitGroup
	wait.Add(1)

	go func() {
		fmt.Println("Hello WG World")
		wait.Done()
	}()

	time.Sleep(time.Second)
}
