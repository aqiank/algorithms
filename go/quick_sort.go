package main

func quickSort(items []int) {
	if len(items) <= 1 {
		return
	}

	medianIndex := len(items) / 2

	items[0], items[medianIndex] = items[medianIndex], items[0]
	medianIndex = 0

	for i := 1; i < len(items); i++ {
		if items[i] < items[medianIndex] {
			items[i], items[medianIndex] = items[medianIndex], items[i]
			medianIndex++
		}
	}

	quickSort(items[:medianIndex])
	quickSort(items[medianIndex+1:])
}
