package main

import (
	"fmt"
)

func startThread(i int) {
	fmt.Println(i)
}

func threads(guard chan struct{}) {
	for i := 0; i < 100; i++ {
		guard <- struct{}{} // would block if guard channel is already filled
		go func(n int) {
			startThread(n)
			<-guard
		}(i)
	}
}

func main() {
	var (
		maxGoroutines = 5
		guard         = make(chan struct{}, maxGoroutines)
	)
	threads(guard)
}
