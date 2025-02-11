package structs

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Address
}

func RunStudentStruct() {
	student1 := Student{
		Name: "Anagh",
		Age:  30,
		Address: Address{
			Street: "99999 Imaginary Street",
			City:   "Victoria",
		},
	}

	fmt.Println("Anagh's current location: ", student1.City)
	fmt.Println("student1: ", student1)

	jsonData, _ := json.Marshal(student1)
	fmt.Println(string(jsonData))
}
