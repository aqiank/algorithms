package main

func binarySearch(items []int, item int) int {
	index := len(items) / 2

	for {
		if item < items[index] {
			index = index / 2
		} else {
			if item == items[index] {
				return index
			}
			index += (len(items) - index) / 2
		}
	}
}
