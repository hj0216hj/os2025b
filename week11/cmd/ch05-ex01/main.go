package main

import "fmt"

func main() {
	subjects := [4]string{"Go", "Javascript", "Python", "Linux"} // slice literal
	subjectSlice := subjects[1:3]                                // slicing
	for _, subject := range subjects {
		fmt.Println(subject)
	}
	fmt.Println("===========================")
	for i := 0; i < len(subjectSlice); i++ {
		fmt.Println(subjectSlice[i])
	}
}
