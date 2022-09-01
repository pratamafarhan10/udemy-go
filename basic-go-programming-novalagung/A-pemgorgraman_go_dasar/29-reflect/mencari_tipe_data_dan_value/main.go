package main

import (
	"fmt"
	"reflect"
)

func main() {
	number := 42
	word := "vedo"

	reflectValueNumber := reflect.ValueOf(number)
	reflectValueString := reflect.ValueOf(word)

	fmt.Println("================ Number")
	fmt.Println("Value of number:", reflectValueNumber.Int())
	fmt.Println("type of number:", reflectValueNumber.Type())
	fmt.Println("Kind:", reflectValueNumber.Kind())
	fmt.Println("CanInt:", reflectValueNumber.CanInt())
	fmt.Println("Interface:", reflectValueNumber.Interface())
	// fmt.Println("String:", reflectValueNumber.String()) it's not a string

	if reflectValueNumber.Kind() == reflect.Int { // cek apakah type data (menggunakan kind) merupakan int atau bukan
		fmt.Println("nilai variabel :", reflectValueNumber.Int())
	}

	fmt.Println("================ String")
	fmt.Println("Value of string:", reflectValueString.String())
	fmt.Println("type of string:", reflectValueString.Type())
	fmt.Println("Kind:", reflectValueString.Kind())
	fmt.Println("CanInt:", reflectValueString.CanInt())
	fmt.Println("Interface:", reflectValueString.Interface())
	// fmt.Println("String:", reflectValueString.String()) it's not a string

	if reflectValueString.Kind() == reflect.String {
		fmt.Println("nilai variabel :", reflectValueString.String())
	}
}
