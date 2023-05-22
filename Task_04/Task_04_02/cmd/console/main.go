package main

import (
	"fmt"

	"example.com/tykoon/pkg/simplebinder"
)

func main() {

	type Product struct {
		Id       int32
		Name     string
		Category string `name:"cat"`
		price    float64
	}

	type Employee struct {
		Id          uint
		FullName    string `name:"name"`
		Departament string `name:"dep"`
		IsActive    bool
	}

	TestProduct := &Product{}
	productData := map[string]string{
		"Name":  "Laptop",
		"cat":   "Hi-Tech",
		"Id":    "34",
		"price": "1234.22",
	}

	TestEmployee := &Employee{}
	employeeData := map[string]string{
		"name":     "John Doe",
		"dep":      "QA-Auto",
		"Id":       "32434",
		"IsActive": "true",
	}

	simplebinder.Bind(productData, TestProduct)
	simplebinder.Bind(employeeData, TestEmployee)

	fmt.Println(*TestProduct)
	fmt.Println(*TestEmployee)
}
