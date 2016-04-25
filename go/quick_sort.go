package main

func quick_sort(items []int) {
    if len(items) <= 1 {
        return
    }

    medianIndex := len(items) / 2

    items[0], items[medianIndex] = items[medianIndex], items[0]
    medianIndex = 0

    for i := 1; i < len(items); i++ {
        if items[i] < items[medianIndex] {
            if i > medianIndex + 1 {
                items[medianIndex], items[medianIndex + 1] = items[medianIndex], items[medianIndex + 1]
            }
            items[i], items[medianIndex] = items[medianIndex], items[i]
            medianIndex++
        }
    }

    quick_sort(items[:medianIndex])
    quick_sort(items[medianIndex + 1:])
}
