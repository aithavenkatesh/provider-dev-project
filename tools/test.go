package main

import "fmt"

func main() {
	data := []int{1, 2, 3}
	fmt.Println("Original Input", data)
	func() {
		for _, v := range data {
			fmt.Println(v)
		}
	}() // Self-contained
}
