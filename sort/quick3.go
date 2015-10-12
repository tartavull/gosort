package sort

func Quick3(array Ints) Ints {

	_quick(array, 0, array.Len()-1)

	return array
}

func _quick3(array Ints, left int, right int) {

	if left >= right {
		return
	}

	lt, rt := partion3(array, left, right)
	_quick3(array, left, lt)
	_quick3(array, rt, right)

}

func partion3(array Ints, left int, right int) (int, int) {

	pivot := left
	l := left + 1

	lt := left
	rt := right
	for ; l <= right; l++ {

		if array.Less(l, pivot) {
			array.Swap(lt, l)
			lt++
		}
		if array.More(l, pivot) {
			array.Swap(l, rt)
			rt--
		}
	}

	return lt, rt
}
