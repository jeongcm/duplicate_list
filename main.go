package main

import (
	"duplicated_list/list"
	"fmt"
	"math/rand"
)

func main() {
	var numbers []int
	// 난수 생성
	for i := 0; i < 5000; i++ {
		num := rand.Intn(50000)
		if num == 0 {
			continue
		}

		numbers = append(numbers, num)
	}

	fmt.Print("duplicate numbers: [")
	result := list.GetDuplicatedNumbers(numbers)
	l := len(result)
	for i, v := range result {
		fmt.Print(v)
		if i == l-1 {
			continue
		}
		fmt.Print(", ")
	}
	fmt.Println("]")
}
