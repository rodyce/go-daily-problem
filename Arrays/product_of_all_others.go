package main

import "fmt"

func productOfAllOthersNaive(L []int) []int {
	ret := make([]int, len(L))
	for i := range ret {
		prod := 1
		for k, q := range L {
			if i == k {
				continue
			}
			prod *= q
		}
		ret[i] = prod
	}
	return ret
}

func productOfAllOthersDivision(L []int) []int {
	ret := make([]int, len(L))
	allProduct := 1
	for _, p := range L {
		allProduct *= p
	}
	for i := 0; i < len(L); i++ {
		ret[i] = allProduct / L[i]
	}
	return ret
}

// Iterator : this iterator struct does not have to be exported, but can be if you have some
// generic Iterator interface that you want/need to implement.
type Iterator struct {
	items []int
	index int
}

func reverseIterator(lst []int) Iterator {
	idx := len(lst)
	return Iterator{lst, idx}
}

// Next returns the next item in the collection.
func (i *Iterator) Next() (int, int) {
	i.index--
	return i.index, i.items[i.index]
}

// HasNext return true if there are values to be read.
func (i *Iterator) HasNext() bool {
	return i.index > 0
}

func productOfAllOthersNoDivision(L []int) []int {
	ret := make([]int, len(L))
	prefixProducts := make([]int, len(L))
	acum := 1
	for i, p := range L {
		acum *= p
		prefixProducts[i] = acum
	}
	suffixProducts := make([]int, len(L))
	acum = 1
	for it := reverseIterator(L); it.HasNext(); {
		i, p := it.Next()
		acum *= p
		suffixProducts[i] = acum
	}

	for i := range ret {
		if i == 0 {
			ret[i] = suffixProducts[i+1]
		} else if i == len(ret)-1 {
			ret[i] = prefixProducts[i-1]
		} else {
			ret[i] = prefixProducts[i-1] * suffixProducts[i+1]
		}
	}

	return ret
}

func main() {
	arr := []int{6, 4, 3}
	res := productOfAllOthersNoDivision(arr)
	fmt.Println(res)
}
