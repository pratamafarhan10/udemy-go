package main

import "fmt"

func main() {
	usa := struct {
		countryName string
		states      []string
		details     map[string]string
	}{
		countryName: "USA",
		states:      []string{"Alabama", "Arkansas", "Arizona"},
		details: map[string]string{
			"abbreviation": "United States of America",
			"continent":    "America",
		},
	}

	fmt.Println(usa)
	fmt.Println(usa.countryName)
	fmt.Println(usa.states, usa.states[0])
	fmt.Println(usa.details, usa.details["abbreviation"])
}
