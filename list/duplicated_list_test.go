package list

import (
	"fmt"
	"testing"
)

func TestGetDuplicateNumber(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 4, 5, 6, 7, 5, 34, 4, 3, 43, 2, 3, 4, 3, 5, 445, 1, 3, 2323, 412, 32}
	n := GetDuplicateNumber(numbers)

	fmt.Println(n)
}
