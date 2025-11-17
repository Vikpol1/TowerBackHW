package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type ConcurMap struct {
	mu  sync.RWMutex
	arr map[string]int
}

func NewMap() *ConcurMap {
	return &ConcurMap{
		arr: make(map[string]int),
	}
}

func (cm *ConcurMap) Set(key string, x int) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.arr[key] = x
}

func (cm *ConcurMap) Get(key string) (int, bool) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	x, exists := cm.arr[key]
	return x, exists
}

func (cm *ConcurMap) Delete(key string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	delete(cm.arr, key)
}

func (cm *ConcurMap) Len() int {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return len(cm.arr)
}

func (cm *ConcurMap) Keys() []string {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	keys := make([]string, 0, len(cm.arr))
	for k := range cm.arr {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	cm := NewMap()
	wg := &sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				key := fmt.Sprintf("w%d-%d", id, j)
				x := rand.Intn(100)
				cm.Set(key, x)
				fmt.Printf("Воркер %d записал %s = %d\n", id, key, x)
			}
		}(i)
	}
	wg.Wait()
	fmt.Printf("\nИтоговый размер мапы: %d\n", cm.Len())
	fmt.Println("Ключи:", cm.Keys())
	check := "w1-1"
	if x, exists := cm.Get(check); exists {
		fmt.Printf("Ключ %s = %d\n", check, x)
	} else {
		fmt.Printf("Ключ %s не найден\n", check)
	}
	fmt.Println("\nВсе значения в мапе:")
	keys := cm.Keys()
	for _, key := range keys {
		if x, exists := cm.Get(key); exists {
			fmt.Printf("  %s = %d\n", key, x)
		}
	}
}
