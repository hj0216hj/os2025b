// average calculates the average of several numbers.
package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
)

func main() {
	weight, err := datafile.GetFloats("meatWeight.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0.0
	for _, number := range weight {
		sum += number
	}
	sampleCount := float64(len(weight))
	fmt.Printf("평균 고기 소비량: %0.2f\n", sum/sampleCount)
}
