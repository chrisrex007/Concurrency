//go:build exp3

package main

import (
	"fmt"
	"sync"
)

var x = 0

func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x++
	m.Unlock()
	wg.Done()
}

func main() {
	wg := new(sync.WaitGroup)
	m := new(sync.Mutex)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(wg,m)
	}

	wg.Wait()
	fmt.Println("Value of x", x)
}
