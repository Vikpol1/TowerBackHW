package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Нужно ввести количество воркеров!")
	}
	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		log.Fatal("Введите положительное число")
	}
	fmt.Printf("Запуск %d воркеров. Нажмите Ctrl+C для остановки.\n", numWorkers)
	rand.Seed(time.Now().UnixNano())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	numCh := make(chan interface{})
	wg := &sync.WaitGroup{}
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, numCh, wg)
	}
	wg.Add(1)
	go generator(ctx, numCh, wg)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	fmt.Println("\nПолучен сигнал прерывания")
	cancel()
	wg.Wait()
	fmt.Println("Все воркеры остановлены. Программа завершена.")
}

func generator(ctx context.Context, numCh chan<- interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(numCh)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			num := rand.Intn(100000)
			fmt.Printf("Продюсер: отправил число %v\n", num)
			select {
			case numCh <- num:
			case <-ctx.Done():
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func worker(ctx context.Context, id int, numCh <-chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Воркер %d: запущен\n", id)
	for {
		select {
		case <-ctx.Done():
			return
		case num, ok := <-numCh:
			if !ok {
				fmt.Printf("Воркер %d: канал закрыт\n", id)
				return
			}
			fmt.Printf("Воркер %d: обработал число %v\n", id, num)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
