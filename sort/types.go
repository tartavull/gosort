package sort

// Ints implements sort.Interface for []int
type Ints []int

func (a Ints) Len() int           { return len(a) }
func (a Ints) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Ints) Less(i, j int) bool { return a[i] < a[j] }
func (a Ints) More(i, j int) bool { return a[i] > a[j] }
