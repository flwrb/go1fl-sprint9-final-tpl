package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	elements := make([]int, size)
	for i := range elements {
		elements[i] = rand.Int()
	}
	return elements
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	maxValue := data[0]
	for _, value := range data {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	var wg sync.WaitGroup
	if len(data) < CHUNKS {
		return maximum(data)
	}
	maxima := make([]int, CHUNKS)
	sizeChunk := len(data) / CHUNKS

	for i := 0; i < CHUNKS; i++ {
		var chunk []int
		if i == CHUNKS-1 {
			chunk = data[i*sizeChunk : len(data)]
		} else {
			chunk = data[i*sizeChunk : i*sizeChunk+sizeChunk]
		}
		wg.Add(1)
		go func(chunk []int, idx int) {
			defer wg.Done()
			maxima[idx] = maximum(chunk)
		}(chunk, i)
	}
	wg.Wait()
	return maximum(maxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	numbers := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(numbers)
	elapsed := time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed.Microseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(numbers)
	elapsed = time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed.Microseconds())
}
