package main

import (
	"C"
	"fmt"
	"sync"
)

//export Goroutines
func Goroutines() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			count := 0
			for j := 1; j <= 20000000; j++ {
				count++
			}
			fmt.Printf("Goroutine finished counting to %d\n", count)
		}()
	}
	wg.Wait()
}

func main() {}
