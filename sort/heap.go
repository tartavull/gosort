package sort

func parent(pos int) int {
	return (pos+1)/2 - 1
}

func children(array Ints, pos int) (child Ints) {

	if a := (pos+1)*2 - 1; array.Len() > a {
		child = append(child, a)
	}

	if b := (pos + 1) * 2; array.Len() > b {
		child = append(child, b)
	}

	return
}

func minChild(array Ints, pos int) int {

	childs := children(array, pos)
	if len(childs) == 0 {
		return -1
	}

	min_pos := childs[0]
	for _, child := range childs {

		if array.Less(child, min_pos) {
			min_pos = child
		}
	}

	return min_pos
}

func swim(array Ints, pos int) {

}

func sink(array Ints, pos int) {

	min_child := minChild(array, pos)
	if min_child == -1 {
		return
	}

	if array.Less(min_child, pos) {
		array.Swap(min_child, pos)
		sink(array, min_child)
	}

}

func Heap(array Ints) Ints {

	for pos := array.Len() / 2; pos >= 0; pos-- {
		sink(array, pos)
	}

	var aux Ints

	for array.Len() > 0 {

		aux = append(aux, array[0])
		array.Swap(0, array.Len()-1)
		array = array[:array.Len()-1]

		sink(array, 0)
	}

	return aux

}
