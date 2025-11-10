package main

import "fmt"

func main() {

	var subject []string
	subject = make([]string, 3)
	subject[0] = "Go"
	subject[2] = "Python"

	for _, subject := range subject {
		fmt.Println(subject)
	}
}
