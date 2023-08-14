package main

import (
	"bytes"
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var (
	workers = 10
	users   = 1000000
)

func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs {
		fmt.Println("Worker", id, "process job", j)
		processData(j)
	}
}

func processData(string) {
	// simulasi proses data
	time.Sleep(1 * time.Second)
}

// goroutine id
func goroutineId() (int, error) {
	goroutinePref := []byte("goroutine ")
	errMessage := errors.New("invalid runtime stuck output")

	buf := make([]byte, 32)
	n := runtime.Stack(buf, false)
	buf = buf[:n]

	buf, ok := bytes.CutPrefix(buf, goroutinePref)
	if !ok {
		return 0, errMessage
	}

	i := bytes.IndexByte(buf, ' ')
	if i < 0 {
		return 0, errMessage
	}

	return strconv.Atoi(string(buf[:i]))
}
