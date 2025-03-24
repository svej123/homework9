package main

import "fmt"

func BS(array []int, needfull int) int {
	left := 0
	right := len(array) - 1

	for left <= right {
		mid := (left + right) / 2
		if array[mid] == needfull {
			return mid
		} else if array[mid] > needfull {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	array := []int{14, 20, 25, 42}
	A := 42
	index := BS(array, A)

	if index != -1 {
		fmt.Printf("Число %d найдено в массиве на индексе %d (порядковый номер: %d)\n", A, index, index+1)
	} else {
		fmt.Println("Число не найдено")
	}
}
