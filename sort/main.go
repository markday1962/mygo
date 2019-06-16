package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

//runs when ever we print
func (p Person) String() string {
	return fmt.Sprintf("%s : %d", p.Name, p.Age)
}

func main() {
	fmt.Println("Running simple sort")
	simpleSort()
	fmt.Println("\nRunning slice sort")
	sliceSort()
}

//ByAge and ByName implements sort.Interface for people based on the
//Age and Name fields, and require a Len, Swap and Less function attached
// to the type
//https://godoc.org/sort#Interface
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []Person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func simpleSort() {
	xi := []int{4, 7, 27, 71, 2, 0, 56, 33, 99}
	xs := []string{"James", "Q", "M", "MoneyPenny", "DR No", "Blofield"}

	fmt.Println("Printing a slice of int", xi)
	sort.Ints(xi)
	fmt.Println("Printing a sorted slice of int", xi)

	fmt.Println("\nPrinting a slice of strings", xs)
	sort.Strings(xs)
	fmt.Println("Printing a sorted slice of strings", xs)
}

func sliceSort() {
	p1 := Person{Name: "James Bond", Age: 47}
	p2 := Person{Name: "Miss Moneypenny", Age: 27}
	p3 := Person{Name: "Dr. No", Age: 70}
	p4 := Person{Name: "M", Age: 52}
	p5 := Person{Name: "Q", Age: 60}

	people := []Person{p1, p2, p3, p4, p5}
	fmt.Println("\nUnsorted slice of people", people)
	sort.Sort(ByAge(people))
	fmt.Println("\nSorted slice of people by age", people)

	sort.Sort(ByName(people))
	fmt.Println("\nSorted slice of people by name", people)

}
