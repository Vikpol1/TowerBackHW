package main

import (
	"fmt"
	"sync"
)

func square(x int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	res := x * x
	out <- res
}
func main() {
	mas := []int{2, 4, 6, 8, 10}
	length := len(mas)
	sum := 0
	resalt := make(chan int, length)
	wg := &sync.WaitGroup{}
	for i := 0; i < length; i++ {
		wg.Add(1)
		go square(mas[i], resalt, wg)
	}
	wg.Wait()
	for i := 0; i < length; i++ {
		sum += <-resalt
	}
	fmt.Println("Сумма квадратов:", sum)
}
