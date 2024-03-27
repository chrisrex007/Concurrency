//

package main

import (
	"fmt"
	"time"
)

func winner() {
	fmt.Println("Winner")
}

func welcome() {
	fmt.Println("Welcome")
}

func main() {
	go winner()
	go welcome()
	time.Sleep(time.Second)
}

// output:
// Welcome
// Winner