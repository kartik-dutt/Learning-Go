package main

import (
	"fmt"
	"sort"
)

/**
 * To create a custom sort, we need the following methods as interface of sort has them.
 * Len() int
 * Less(i, j int) bool
 * Swap(i, j int)
 */

// Odd even sort.
type OddEven []int

func (arr OddEven) Len() int {
	return len(arr)
}

func (arr OddEven) Less(i, j int) bool {
	if arr[i]&1 == 1 {
		if arr[j]&1 == 1 {
			return arr[i] < arr[j]
		} else {
			return true
		}
	}

	return arr[i] < arr[j]
}

func (arr OddEven) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

type person struct {
	name string
	age  int
}

func main() {
	// 1.) Sort integers and check if they are sorted.
	vals := []int{9, 8, 1, 4, 5, 2, 24}
	print := fmt.Println
	print("Before : ", vals, sort.IntsAreSorted(vals))
	sort.Ints(vals)
	print("After : ", vals, sort.IntsAreSorted(vals))
	// Similary for other data types.
	a := []string{"def", "abc", "ghi"}
	print("Before :", a)
	sort.Strings(a)
	print("After :", a)

	// 2.) Custom sorting.
	sort.Sort(OddEven(vals))
	print("Odd Even Sort", vals)

	// Let's reset.
	sort.Ints(vals)
	// 3.) Custom Sorting made simple.
	sort.Slice(vals, func(i, j int) bool {
		if vals[i]&1 == 1 {
			if vals[j]&1 == 1 {
				return vals[i] < vals[j]
			} else {
				return true
			}
		}

		return vals[i] < vals[j]
	})

	print("Inline Odd Even Sort", vals)

	persons := []person{person{name: "k", age: 21}, person{name: " q", age: 18}}
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].age < persons[j].age
	})
	fmt.Println(persons)
}
