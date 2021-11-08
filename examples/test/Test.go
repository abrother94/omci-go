package main

import (
	"fmt"
)

func init() {
	fmt.Println("init 2")
}

func init() {
	fmt.Println("init 1")
}

func main() {
	fmt.Println("Hello, playground")
}
