package main

import (
    "testing"
)

func TestInsertionSort(t *testing.T) {
    sample := []int{ 378, 131, 378, 572, 712, 178, 751, 293, 578, 423, 188, 51, 137, 659 }
    desired_result := []int{ 51, 131, 137, 178, 188, 293, 378, 378, 423, 572, 578, 659, 712, 751 }

    insertion_sort(sample)

    for i := range sample {
        if sample[i] != desired_result[i] {
            t.Error("insertion_sort")
            return
        }
    }
}
