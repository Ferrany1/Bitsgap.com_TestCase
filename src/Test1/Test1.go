package Test1

import (
	"fmt"
	"sync"
)

func printVal(j int) {
	fmt.Print(j)
}

func group(wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 1; j <= 3; j++ {
		printVal(j)
	}
	fmt.Println()
}

func Test1() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go group(&wg)
	}
	wg.Wait()
}
