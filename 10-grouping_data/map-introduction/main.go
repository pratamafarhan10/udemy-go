package main

import "fmt"

func main() {
	// Map allow us to store a key value pair
	// map[key data type]value data type {data with key value pair}
	peopleAge := map[string]int{
		"farhan": 22,
		"james":  42,
		"sule":   30,
	}

	fmt.Println(peopleAge)

	fmt.Println("farhan age", peopleAge["farhan"])

	dumbledore := peopleAge["dumbledore"]

	fmt.Println("dumbledore age is", dumbledore) // if the key doesn't stored in the map, the value will be the zero value of the data type. (THIS CAN BE A PROBLEM)

	// Check wether the key is stored or nah
	v, ok := peopleAge["dumbledore"]

	fmt.Println("check the key stored in the map with the v ok idiom:", v, ok)

	// we can check using an if statement
	if v, ok := peopleAge["farhan"]; ok {
		fmt.Println("THIS WILL PRINT", v)
	}

	buahBatu := [][]string{
		{"farhan", "pratama"},
		{"james", "bond"},
	}

	bojongsoang := [][]string{
		{"miss", "moneypenny"},
		{"andre", "taulany"},
	}

	kecamatan := map[string][][]string{
		"buah batu":   buahBatu,
		"bojongsoang": bojongsoang,
	}
	fmt.Println("================ A map of multidimensional slice")
	for i, val := range kecamatan {
		fmt.Println(i, val)
	}

	kabupaten := map[string]map[string][][]string{
		"bandung": kecamatan,
	}

	fmt.Println("================ A map of a map of multidimensional slice")
	for i, val := range kabupaten {
		fmt.Println(i, val)
	}
}
