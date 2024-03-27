//go:build exp2

package main

import (
	"fmt"
	"sync"
)

func winner(wg *sync.WaitGroup) {
	defer wg.Done() //decrement 1 in counter
	fmt.Println("Winner")
}

func welcome(wg *sync.WaitGroup) {
	defer wg.Done() //decrement 1 in counter
	fmt.Println("Welcome")
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2) //adds 2 to the counter
	go winner(wg)
	go welcome(wg)
	wg.Wait()
}
