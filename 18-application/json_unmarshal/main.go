package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       int    `json:"Age"`
}

type secretAgent struct {
	person
	deployedTo string
}

func main() {
	fmt.Println("============= Struct marshal")
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

	byteSecretAgents, err := json.Marshal(secretAgents)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(byteSecretAgents))

	fmt.Println("============= Struct unmarshal")

	var secretAgentsConverted []secretAgent

	err = json.Unmarshal(byteSecretAgents, &secretAgentsConverted)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(secretAgentsConverted)

	fmt.Println("============= Map of slice of string marshal")

	province := map[string][]string{
		"indonesia": []string{"Jakarta", "Jawa Barat", "Jawa Tengah"},
		"usa":       []string{"Alabama", "Los Angelese", "Texas"},
	}

	byteProvince, err := json.Marshal(province)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(byteProvince))

	fmt.Println("============= Map of slice of string unmarshal")

	var provinceConverted map[string][]string

	err = json.Unmarshal(byteProvince, &provinceConverted)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(provinceConverted)

	fmt.Println("============= Map of struct marshal")
	secretAgents2 := map[string]secretAgent{
		"james": sa1,
		"bean":  sa2,
	}

	byteSecretAgents2, err := json.Marshal(secretAgents2)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(byteSecretAgents2))

	fmt.Println("============= Map of struct unmarshal")

	var secretAgents2Converted map[string]secretAgent

	err = json.Unmarshal(byteSecretAgents2, &secretAgents2Converted)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(secretAgents2Converted)

	fmt.Println("============= Unmarshal")

	request := `[{"FirstName":"james","LastName":"bond","Age":42},{"FirstName":"mr","LastName":"bean","Age":53}]`

	byteRequest := []byte(request)

	var convertedRequest []secretAgent

	err = json.Unmarshal(byteRequest, &convertedRequest)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(convertedRequest)
}
