package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

func main() {

	n := 10
	numbers := make([]int, n)

	ch := make(chan int)
	resCh := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range n {
			num := rand.IntN(101)
			numbers[i] = num
			ch <- num
		}
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch {
			resCh <- num * num
		}
		close(resCh)
	}()

	go func() {
		wg.Wait()
	}()

	for result := range resCh {
		fmt.Println(result)
	}
}
