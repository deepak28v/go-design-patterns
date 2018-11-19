package concurrency

import (
	"fmt"
	"time"
)

func anonymousFunc() {
	messagePrinter := func(msg string) {
		fmt.Println(msg)
	}

	go messagePrinter("Hello World 1")
	go messagePrinter("Hello World 2")
	time.Sleep(time.Second)
}
