package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func BenchmarkWorkerOverGoroutine(b *testing.B) {
	fmt.Println("example go worker")

	jobs := make(chan string, users)
	var wg sync.WaitGroup
	for i := 1; i < users; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
		jobs <- fmt.Sprintf("user %d", i)
		fmt.Printf("goroutine berjalan %d || ", runtime.NumGoroutine())
		time.Sleep(1 * time.Second)
	}
	close(jobs)
	wg.Wait()
}

func BenchmarkWorkerPool(b *testing.B) {
	fmt.Println("example go worker")

	jobs := make(chan string, users)
	var wg sync.WaitGroup

	for i := 1; i < 5; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
		fmt.Printf("open goroutine %d\n", runtime.NumGoroutine())
		time.Sleep(1 * time.Second)
	}
	for i := 1; i < users; i++ {
		jobs <- fmt.Sprintf("user %d", i)
	}
	close(jobs)
	wg.Wait()

	time.Sleep(10 * time.Second)
}
