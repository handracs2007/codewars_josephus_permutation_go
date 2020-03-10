package main

import "fmt"

func Josephus(items []interface{}, k int) []interface{} {
	result := make([]interface{}, 0)
	currentIndex := 0

	for len(items) > 0 {
		currentIndex = (currentIndex + k - 1) % len(items)
		item := items[currentIndex]
		tempItems := append(items[0:currentIndex])

		if currentIndex+1 < len(items) {
			tempItems = append(tempItems, items[currentIndex+1:]...)
		}

		items = tempItems
		result = append(result, item)
	}

	return result
}

func main() {
	items := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//result := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	k := 3
	fmt.Println(Josephus(items, k))
}
