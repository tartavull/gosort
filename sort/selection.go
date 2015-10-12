package sort

func Selection(array Ints) Ints {

	for i := 0; i < array.Len(); i++ {

		min_pos := i
		for j := i; j < array.Len(); j++ {

			if array.Less(j, min_pos) {
				min_pos = j
			}

		}
		array.Swap(i, min_pos)

	}

	return array
}
