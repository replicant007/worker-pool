package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
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
			fmt.Printf("No int elements left for worker %d to process\n", id)
			return
		}

		number := (*nums)[0]
		*nums = (*nums)[1:]
		mu.Unlock()

		fmt.Printf("Worker %d processed data %d -> result %d\n", id, number, number*number)
		delay := time.Duration(rand.Intn(10))
		time.Sleep(delay * time.Second)

		updateMetrics(id, delay)
	}
}

func channelWorker(id int, ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var delay time.Duration

	for letter := range ch {
		fmt.Printf("Worker %d processed data %s -> result %s\n", id, letter, strings.ToUpper(strings.Repeat(letter, 2)))
		delay = time.Duration(rand.Intn(10))
		time.Sleep(delay * time.Second)
		updateMetrics(id, delay)
	}
	fmt.Printf("No str elements left for worker %d to process\n", id)
}

func updateMetrics(id int, delay time.Duration) {
	if value, exists := metrics.Load(id); !exists {
		metrics.Store(id, workerMetrics{1, delay})
	} else {
		wm := value.(workerMetrics)
		wm.taskCount += 1
		wm.totalTime += delay
		metrics.Store(id, wm)
	}
}

func main() {

	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	ch := make(chan string, len(letters))
	var wg sync.WaitGroup

	for i := 0; i < len(letters); i++ {
		ch <- letters[i]
	}
	close(ch)

	for i := 0; i < 4; i++ {
		wg.Add(2)
		go worker(i, &nums, &wg)
		go channelWorker(i+4, ch, &wg)
	}

	wg.Wait()
	fmt.Println("\nAll workers have finished processing.")

	type element struct {
		key   int
		value workerMetrics
	}

	var elements []element
	metrics.Range(func(key, value any) bool {
		elements = append(elements, element{key.(int), value.(workerMetrics)})
		return true
	})
	sort.Slice(elements, func(i, j int) bool {
		return elements[i].key < elements[j].key
	})
	for _, el := range elements {
		fmt.Printf("Worker %d → Tasks: %d, Total Time: %v\n", el.key, el.value.taskCount, el.value.totalTime)
	}
}
