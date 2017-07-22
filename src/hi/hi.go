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
