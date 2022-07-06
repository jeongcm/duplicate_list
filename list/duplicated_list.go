package list

import (
	"errors"
	"fmt"
)

func GetDuplicateNumber(numbers []int) []string {
	// numberMap 은 입력받은 숫자 배열을 map으로 변환하고 중복된 횟수만큼 value 값을 높인다.
	numberMap := make(map[int]int)
	// result 는 중복된 숫자와 1회 이상의 중복 횟수를 문자열로 변환해서 담는다.
	var result []string
	for _, number := range numbers {
		if number < 1 || number > 50000 {
			fmt.Println(errors.New("invalid number"))
			return nil
		}
		numberMap[number]++
	}

	// numberMap 에 저장된 중복된 숫자와 횟수를 확인하고 중복이 2개이상 됐을 경우 결과에 값을 저장한다.
	for key, value := range numberMap {
		// 같은 숫자를 최소 2개 이상 가지고 있는 경우에만 결과값에 추가한다.
		if value < 2 {
			continue
		}
		// value 로 가진 값에서 1을 빼서 중복된 횟수를 확인한다.
		result = append(result, fmt.Sprintf("%d:%d", key, value-1))
	}

	return result
}
