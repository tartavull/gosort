package main

import (
	"fmt"
	"github.com/tartavull/gosort/sort"
)

func copy_slice(original []int) []int {

	copied := make([]int, len(original))
	copy(copied, original)

	return copied
}

func main() {

	array := []int{9, 3, 4, 2, 1, 10, 2}

	fmt.Printf(" original %v \n", array)

	fmt.Printf("selection %v \n", sort.Selection(copy_slice(array)))
	fmt.Printf("insertion %v \n", sort.Insertion(copy_slice(array)))
	fmt.Printf("shell %v \n", sort.Shell(copy_slice(array)))
	fmt.Printf("merge %v \n", sort.Merge(copy_slice(array)))
	fmt.Printf("quick %v \n", sort.Quick(copy_slice(array)))
	fmt.Printf("quick3 %v \n", sort.Quick3(copy_slice(array)))
	fmt.Printf("heap %v \n", sort.Heap(copy_slice(array)))

}
