package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Friends []Person `json:"friends`
}

func main() {
	friend := Person{Name: "X", Age: 100, Friends: []Person{}}
	person := Person{Name: "Kartik", Age: 21, Friends: []Person{friend}}
	fmt.Println(person)
	byteArray, _ := json.Marshal(person)
	fmt.Println(string(byteArray))

	// Pretty Print
	byteArray, _ = json.MarshalIndent(person, "", "  ")
	fmt.Println(string(byteArray))

	// Json to struct.
	var samePerson Person
	_ = json.Unmarshal(byteArray, &samePerson)
	fmt.Println(samePerson)

	// From strings
	str := []byte(`{"name":"Kartik","age":21,"Friends":[{"name":"X","age":100,"Friends":[]}]}`)
	_ = json.Unmarshal(str, &samePerson)
	fmt.Println(samePerson)

}
