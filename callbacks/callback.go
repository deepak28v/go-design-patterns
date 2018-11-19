package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	var wait sync.WaitGroup

	wait.Add(1)
	toUpperSync("Hello Callbacks !!", func(v string) {
		fmt.Printf("Callback: %s\n", v)
		wait.Add(-1)
	})

	println("Waiting async response...")
	//wait.Wait()
}

func toUpperSync(word string, f func(string)) {
	go func() {
		f(strings.ToUpper(word))
	}()
}
