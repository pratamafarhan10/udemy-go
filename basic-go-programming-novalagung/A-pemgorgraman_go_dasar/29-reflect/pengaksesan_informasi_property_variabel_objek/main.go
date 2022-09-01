package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Grade int
}

func (s *Student) getPropertyInfo() {
	reflectValue := reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Pointer {
		reflectValue = reflectValue.Elem() // isinya value adalah si struct student
	}

	reflectType := reflectValue.Type() // NYIMPEN TIPE YAITU STRUCT

	fmt.Println("get field by name (Name):", reflectValue.FieldByName("Name"))
	fmt.Println("get field by name (Grade):", reflectValue.FieldByName("Grade"))

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("============ index number", i)
		fmt.Println("Field name:\t", reflectType.Field(i).Name)      // Check fieldname dari struct
		fmt.Println("Tipe data:\t", reflectType.Field(i).Type)       // check type data dari fieldname struct
		fmt.Println("Nilai:\t\t", reflectValue.Field(i).Interface()) // ngambil value dari struct berdasarkan i
		fmt.Println("")
	}
}

func (s *Student) setName(name string) {
	s.Name = name
}

func main() {
	s1 := Student{
		Name:  "Andre",
		Grade: 90,
	}

	fmt.Println("========== Get property info")
	(&s1).getPropertyInfo()

	fmt.Println("========== Get function info")
	reflectValue := reflect.ValueOf(&s1)

	method := reflectValue.MethodByName("setName")
	// method := reflectValue.Method(1)
	method.Call([]reflect.Value{
		reflect.ValueOf("Wick"),
	})

	fmt.Println("nama :", s1.Name)
}
