package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

type ByAge []Person

type ByName []Person

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	// Because in Go slices are pointers (who also have length and capacity) we don't need to capture what the sort.Ints() func returns.
	// The underlying array of the slice is mutated within the func
	sort.Ints(xi)
	fmt.Println(xi)

	sort.Strings(xs)
	fmt.Println(xs)

	// Custom sorting
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}
	fmt.Println(people)

	// Let's sort this slice by age:
	sort.Sort(ByAge(people))
	fmt.Println(people)

	// Now let's sort it by first name:
	sort.Sort(ByName(people))
	fmt.Println(people)
}

// The following three methods work with ByAge type when the built in sort.Sort() function is invoked
func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

// The following three methods work with ByName type when the built in sort.Sort() function is invoked
func (a ByName) Len() int {
	return len(a)
}

func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByName) Less(i, j int) bool {
	return a[i].First < a[j].First
}
