package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mu sync.Mutex
var metrics sync.Map

type workerMetrics struct {
	taskCount int
	totalTime time.Duration
}

func worker(id int, nums *[]int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		mu.Lock()
		if len(*nums) == 0 {
			mu.Unlock()
			fmt.Printf("No elements left for worker %d to process\n", id)
			return
		}

		x := (*nums)[0]
		*nums = (*nums)[1:]
		mu.Unlock()

		fmt.Printf("Worker %d processed data %d -> result %d\n", id, x, x*x)
		delay := time.Duration(rand.Intn(10))
		time.Sleep(delay * time.Second)

		if value, exists := metrics.Load(id); !exists {
			metrics.Store(id, workerMetrics{1, delay})
		} else {
			wm := value.(workerMetrics)
			wm.taskCount += 1
			wm.totalTime += delay
			metrics.Store(id, wm)
		}
	}
}

func main() {

	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go worker(i, &nums, &wg)
	}

	wg.Wait()
	fmt.Println("All workers have finished processing.")
	fmt.Println()

	metrics.Range(func(key, value any) bool {
		id := key.(int)
		wm := value.(workerMetrics)
		fmt.Printf("Worker %d → Tasks: %d, Total Time: %v\n", id, wm.taskCount, wm.totalTime)
		return true
	})
}
