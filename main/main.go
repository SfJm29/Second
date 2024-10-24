package main

import (
	"fmt"
	"sync"
)

func main() {
	var amountofgoroutines = 5
	var numbergoroutine = 0
	var wg sync.WaitGroup
	wg.Add(amountofgoroutines)
	for i := 1; i <= amountofgoroutines; i++ {
		numbergoroutine++
		go func(ng int) {
			for j := 0; j < 10; j++ {
				fmt.Printf("My number: %v\n", ng)
			}
			wg.Done()
		}(numbergoroutine)
	}
	wg.Wait()
}
