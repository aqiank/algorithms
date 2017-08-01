package main

func mergeLists(lists [][]int) []int {
	for {
		c := [][]int{}

		for i := 0; i < len(lists); i += 2 {
			if i+1 >= len(lists) {
				c = append(c, lists[i])
				continue
			}
			c = append(c, merge(lists[i], lists[i+1]))
		}

		lists = c

		if len(c) == 1 {
			return c[0]
		}
	}
}

func merge(a, b []int) []int {
	c := []int{}
	ia, ib := 0, 0

	for {
		if a[ia] < b[ib] {
			c = append(c, a[ia])
			ia++
			if ia == len(a) {
				c = append(c, b[ib:]...)
				break
			}
		} else {
			c = append(c, b[ib])
			ib++
			if ib == len(b) {
				c = append(c, a[ia:]...)
				break
			}
		}
	}

	return c
}

func divide(items []int) [][]int {
	lists := [][]int{}

	for _, v := range items {
		lists = append(lists, []int{v})
	}

	return lists
}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}

	lists := divide(items)
	return mergeLists(lists)
}
