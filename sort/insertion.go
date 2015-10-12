package sort

func Insertion(array Ints) Ints {

	for i := 1; i < array.Len(); i++ {

		for j := i; j > 0; j-- {

			if array.Less(j, j-1) {

				array.Swap(j, j-1)

			} else {
				break
			}

		}
	}

	return array

}
