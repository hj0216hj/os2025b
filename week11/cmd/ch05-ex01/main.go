package main

import "fmt"

func main() {
	subjects := [4]string{"Go", "Javascript", "Python", "Linux"} // slice literal
	subjectSlice := subjects[:3]
	subjects[0] = "Java"
	//subjectSlice[0] = "Java" // << 슬라이스가 바뀌면 원본도 같이 바뀐다. 같은 곳을 바라보고 있다.
	for _, subject := range subjects {
		fmt.Println(subject)
	}
	fmt.Println("===========================")
	for i := 0; i < len(subjectSlice); i++ {
		fmt.Println(subjectSlice[i])
	}
}
