package main

func swap(items []int, i, j int) {
    tmp := items[i]
    items[i] = items[j]
    items[j] = tmp
}

func quick_sort(items []int) {
    if len(items) <= 1 {
        return
    }

    medianIndex := len(items) / 2
    swap(items, 0, medianIndex)
    medianIndex = 0

    for i := 1; i < len(items); i++ {
        if items[i] < items[medianIndex] {
            if i > medianIndex + 1 {
                swap(items, medianIndex, medianIndex + 1)
            }
            swap(items, medianIndex, i)
            medianIndex++
        }
    }

    quick_sort(items[:medianIndex])
    quick_sort(items[medianIndex + 1:])
}
