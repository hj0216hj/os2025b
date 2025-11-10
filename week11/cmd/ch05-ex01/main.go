package main

import "fmt"

func main() {

	//slice v0.1에서 for에 subject 변수이름을 바꿔야 작동함
	//var subjects []string
	//subjects = make([]string, 3)

	//subjects := make([]string, 3)

	//subjects[0] = "Go"
	//subjects[2] = "Python"

	subjects := []string{"Go", "", "Python"} // slice literal
	for _, subject := range subjects {
		fmt.Println(subject)
	}
}
