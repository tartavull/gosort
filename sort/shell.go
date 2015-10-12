package sort

func Shell(array Ints) Ints {

	steps := [...]int{701, 301, 132, 57, 23, 10, 4, 1}

	for step := range steps {

		for i := step; i < array.Len(); i++ {

			for j := i; j-step >= 0; j -= step {

				if array.Less(j, j-step) {
					array.Swap(j, j-step)
				} else {
					break
				}
			}

		}

	}

	return array
}
