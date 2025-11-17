package main

import (
	"fmt"
	"sync"
)

func main() {
	num := []int{1, 3, 5, 7, 9, 11}
	first := make(chan int, len(num))
	second := make(chan int, len(num))
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(first)
		for _, x := range num {
			first <- x
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(second)
		for x := range first {
			res := x * 2
			second <- res
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for res := range second {
			fmt.Println(res)
		}
	}()
	wg.Wait()
	fmt.Println("Конец!")
}
