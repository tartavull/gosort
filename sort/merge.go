package sort

func merger(array Ints, left int, mid int, right int) {

	aux := make([]int, right-left)

	l := left
	m := mid
	for i, _ := range aux {

		if l == mid {
			aux[i] = array[m]
			m++
		} else if m == right {
			aux[i] = array[l]
			l++
		} else if array.Less(l, m) {
			aux[i] = array[l]
			l++
		} else {
			aux[i] = array[m]
			m++
		}

	}
	copy(array[left:right], aux)
}

func Min(a, b int) int {

	if a < b {
		return a
	}
	return b
}

func Merge(array Ints) Ints {

	for width := 1; width < array.Len(); width *= 2 {

		for bin := 0; bin < array.Len(); bin += 2 * width {

			left := bin
			mid := bin + width
			right := Min(array.Len(), bin+2*width)

			merger(array, left, mid, right)

		}
	}

	return array

}
