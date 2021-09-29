package main

import "fmt"

func failureFunc() {
	panic("I failed!")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Caught error: %q", r)
		}
	}()

	failureFunc()
	fmt.Println("I win!")
}
