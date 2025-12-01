package main

import (
	"fmt"
)

func main() {
	// Last In First Out 스택구조
	defer fmt.Println("1st defer")
	defer fmt.Println("2st defer")
	defer fmt.Println("3st defer")
	fmt.Println("main logic")
}
