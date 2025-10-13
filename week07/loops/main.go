package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// shadowing
	// var int int = 99
	// var b int = 8
	//var fmt string = "inha"
	//fmt.Println(int, b)
	//fmt.Println(fmt)

	fmt.Println("Enter a grade :")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)                // 문자열 주위에 붙은 공간 및 탭 키 등 제거
	score, err := strconv.ParseFloat(i, 64) // 정리된 문자열을 실수타입으로 변환
	var status string
	if score >= 60 {
		status = "passing"
		//status := "passig" << 지역변수 괄호 안에서만 존재
	} else {
		status = "falling"
	}
	fmt.Println(score, status)
}
