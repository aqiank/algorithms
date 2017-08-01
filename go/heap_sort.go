package main

func parent(idx int) int {
	return (idx - 1) / 2
}

func left(idx int) int {
	return 2*idx + 1
}

func right(idx int) int {
	return 2*idx + 2
}

func heapify(items []int, p int) {
	if p < 0 || p >= len(items) {
		return
	}

	l := left(p)
	r := right(p)
	s := -1

	// If we have child at all, compare with parent
	if l < len(items) {
		// Get the better sibling
		if r < len(items) && items[r] > items[l] {
			s = r
		} else {
			s = l
		}

		// Compare the better sibling to parent, swap if better than parent
		if items[s] > items[p] {
			items[p], items[s] = items[s], items[p]

			heapify(items, s)
		}
	}

	heapify(items, p-1)
}

func sift(items []int, max int) {
	for max >= 0 {
		items[0], items[max] = items[max], items[0]
		max--
		heapify(items[:max+1], max)
	}
}

func heapSort(items []int) {
	heapify(items, len(items)-1)
	sift(items, len(items)-1)
}
