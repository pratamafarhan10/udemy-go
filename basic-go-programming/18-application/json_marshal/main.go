package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName, LastName string
	Age                 int
}

type secretAgent struct {
	person
	deployedTo string
}

func main() {
	fmt.Println("============= Struct")
	sa1 := secretAgent{
		person: person{
			FirstName: "james",
			LastName:  "bond",
			Age:       42,
		},
		deployedTo: "Russia",
	}

	sa2 := secretAgent{
		person: person{
			FirstName: "mr",
			LastName:  "bean",
			Age:       53,
		},
		deployedTo: "UK",
	}

	secretAgents := []secretAgent{sa1, sa2}

	fmt.Println(secretAgents)

	b, err := json.Marshal(secretAgents)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(b))

	fmt.Println("============= Map of slice of string marshal")

	province := map[string][]string{
		"indonesia": []string{"Jakarta", "Jawa Barat", "Jawa Tengah"},
		"usa":       []string{"Alabama", "Los Angelese", "Texas"},
	}

	b, err = json.Marshal(province)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(b))

	fmt.Println("============= Map of struct marshal")
	secretAgents2 := map[string]secretAgent{
		"james": sa1,
		"bean":  sa2,
	}

	b, err = json.Marshal(secretAgents2)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(b))
}
