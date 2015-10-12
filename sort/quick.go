package sort

func Quick(array Ints) Ints {

	_quick(array, 0, array.Len()-1)

	return array
}

func _quick(array Ints, left int, right int) {

	if left >= right {
		return
	}

	pivot := partion(array, left, right)
	_quick(array, left, pivot-1)
	_quick(array, pivot+1, right)

}

func partion(array Ints, left int, right int) int {

	pivot := left
	l := left + 1
	r := right
	for {

		for ; array.Less(l, pivot); l++ {
		}
		for ; array.More(r, pivot); r-- {
		}

		if l > r {
			break
		}

		array.Swap(l, r)
		l++
		r--

	}
	array.Swap(left, l-1)

	return l - 1
}
