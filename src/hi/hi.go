package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hi, Lirian")

	defer tryDefer()

	fmt.Println(tryError(true))
	fmt.Println(tryError(false))

	tryGo()
}

func tryDefer() {
	fmt.Println("this is a defer")
}

func tryError(expectError bool) error {
	if expectError {
		return errors.New("new error")
	} else {
		return nil
	}
}

func tryChannel(a int, b chan int) {
	fmt.Printf("Executing %s for %s\n", a, b)
	b <- a
}

func tryGo() {
	b := make(chan int)

	go tryChannel(1, b)
	go tryChannel(2, b)
	go tryChannel(4, b)

	fmt.Println(<-b, <-b, <-b)
}
