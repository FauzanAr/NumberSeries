package main

import "fmt"


func calculate(numbers []int) int {
	var postNum, negNum []int
	for i, item := range numbers {
		if i + 1 < len(numbers) {
			if item + 1 == numbers[i+1] {
				if len(postNum) > 0 {
					postNum = removeIndex(postNum, len(postNum) - 1)
				}
				postNum = append(postNum, item)
				postNum = append(postNum, numbers[i+1])
			}
			if item - 1 == numbers[i+1] {
				if len(negNum) > 0 {
					negNum = removeIndex(negNum, len(negNum) - 1)
				}
				negNum = append(negNum, item)
				negNum = append(negNum, numbers[i+1])
			}
		}
	}
	return max(union(postNum, negNum))
}

func removeIndex(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}

func union(firstArr, secondArr []int) []int {
	m := make(map[int]bool)
	var result []int

	for _, item := range firstArr {
		m[item] = true
	}

	for _, item := range secondArr {
		if _, ok := m[item]; ok {
			result = append(result, item)
		}
	}
	return result
}

func max(numbers []int) int {
	var max int
	for _, item := range numbers {
		if item > max {
			max = item
		}
	}
	return max
}

func main() {
	var arr = []int{7, 1, 2, 9, 7, 2, 1}
	fmt.Println(calculate(arr))
}
