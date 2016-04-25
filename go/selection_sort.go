package main

func selection_sort(items []int) {
    if len(items) < 2 {
        return
    }

    for j := len(items) - 1; j >= 0; j-- {
        maxPos := j
        for i := j - 1; i >= 0 && i < j; i-- {
            if items[i] > items[maxPos] {
                maxPos = i
            }
        }
        if maxPos != j {
            items[j], items[maxPos] = items[maxPos], items[j]
        }
    }
}
