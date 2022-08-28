package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func main() {
	dir := "./"

	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	employees := []User{
		{"John", "21", "23423122", "Cleverinsight", Addres{"Bangalore", "Karnataka", "India", "111111"}},
		{"Poul", "22", "23423122", "Cleverinsight", Addres{"Bangalore", "Karnataka", "India", "222222"}},
		{"Robert", "23", "23423122", "Cleverinsight", Addres{"Bangalore", "Karnataka", "India", "333333"}},
		{"Vince", "24", "23423122", "Cleverinsight", Addres{"Bangalore", "Karnataka", "India", "444444"}},
		{"Albert", "25", "23423122", "Cleverinsight", Addres{"Bangalore", "Karnataka", "India", "555555"}},
	}

	for _, value := range employees {
		db.Write("users", value.Name, User{
			Name:    value.Name,
			Age:     value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(records)

	allusers := []User{}

	for _, f := range records {
		employeeFound := User{}
		if err := json.Unmarshal([]byte(f), &employeeFound); err != nil {
			fmt.Println("Error", err)
		}
		allusers = append(allusers, employeeFound)
	}
}
