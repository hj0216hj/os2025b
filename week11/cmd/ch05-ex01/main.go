package main

import "fmt"

func main() {
	subjects := [4]string{"Go", "Javascript", "Python", "Linux"} // slice literal
	subjectSlice := subjects[:3]
	subjects[0] = "Java"
	subjectSlice = append(subjectSlice, "Go") // append(subjectSlice, "Go", "Database") << 원본에 둘다 추가가 안됨
	for _, subject := range subjects {
		fmt.Println(subject)
	}
	fmt.Println("===========================")
	for i := 0; i < len(subjectSlice); i++ {
		fmt.Println(subjectSlice[i])
	}
}
