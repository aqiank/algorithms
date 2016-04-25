package main

func insertion_sort(items []int) {
    if len(items) < 2 {
        return
    }

    for j := 1; j < len(items); j++ {
        value := items[j]

        i := j - 1
        for i >= 0 && items[i] > value {
            items[i + 1] = items[i]
            i--
        }

        items[i + 1] = value
    }
}
