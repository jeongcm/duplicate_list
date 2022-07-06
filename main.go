package main

import (
	"duplicated_list/list"
	"fmt"
	"math/rand"
)

func getRandomNumbers(maxLength, maxNumber int) []int {
	var numbers []int
	// 난수 생성
	for i := 0; i < maxLength; i++ {
		num := rand.Intn(maxNumber)
		if num == 0 {
			continue
		}

		numbers = append(numbers, num)
	}

	return numbers
}

func printDuplicateNumbers(result []string) {
	l := len(result)

	fmt.Print("duplicate numbers: [")
	for i, v := range result {
		fmt.Print(v)
		if i == l-1 {
			continue
		}
		fmt.Print(", ")
	}
	fmt.Println("]")
}

func main() {
	numbers := getRandomNumbers(5000, 50000)
	result := list.GetDuplicatedNumbers(numbers)

	printDuplicateNumbers(result)
}
