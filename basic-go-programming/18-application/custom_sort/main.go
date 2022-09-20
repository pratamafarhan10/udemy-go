package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{
		name: "andre",
		age:  42,
	}

	p2 := person{
		name: "sule",
		age:  45,
	}

	p3 := person{
		name: "james",
		age:  39,
	}

	people := []person{p1, p2, p3}
	fmt.Println("======= Before sort")
	fmt.Println(people)

	fmt.Println("======= After sort ascending by age")
	sort.Slice(people, func(i, j int) bool { return people[i].age < people[j].age })
	fmt.Println(people)

	fmt.Println("======= After sort ascending by name")
	sort.Slice(people, func(i, j int) bool { return people[i].name < people[j].name })
	fmt.Println(people)

	fmt.Println("======= After sort descending by age")
	sort.Slice(people, func(i, j int) bool { return people[i].age > people[j].age })
	fmt.Println(people)

	fmt.Println("======= After sort descending by name")
	sort.Slice(people, func(i, j int) bool { return people[i].name > people[j].name })
	fmt.Println(people)

}
