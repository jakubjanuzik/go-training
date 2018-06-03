package main

func each(array []int, fn func(number int)) {

	for _, item := range array {
		fn(item)
	}
}

func print(sth int) {
	println(sth)
}

func filter(arr []int, predicate func(int) bool) []int {
	result := make([]int, 0)
	each(arr, func(value int) {
		if predicate(value) {
			result = append(result, value)
		}
	})
	return result
}

func isEven(element int) bool {
	return element%2 == 0
}

func mainEx01() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	each(numbers, print)
	each(filter(numbers, isEven), print)
}
