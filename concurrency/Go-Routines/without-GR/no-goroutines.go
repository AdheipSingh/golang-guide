package main

import (
	"fmt"
)

func main() {
	one()
	two()
}

func one() {
	for i := 0; i < 100; i++ {
		fmt.Println("One", i)
	}
}

func two() {
	for i := 0; i < 100; i++ {
		fmt.Println("Two", i)
	}
}
