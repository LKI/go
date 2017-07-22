package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hi, Lirian")

	defer TryDefer()

	fmt.Println(TryError(true))
	fmt.Println(TryError(false))

	TryGo()
}

func TryDefer() {
	fmt.Println("this is a defer")
}

func TryError(expectError bool) error {
	if expectError {
		return errors.New("new error")
	} else {
		return nil
	}
}

func TryChannel(a int, b chan int) {
	fmt.Printf("Executing %s for %s\n", a, b)
	b <- a
}

func TryGo() {
	b := make(chan int)

	go TryChannel(1, b)
	go TryChannel(2, b)
	go TryChannel(4, b)

	fmt.Println(<-b, <-b, <-b)
}
