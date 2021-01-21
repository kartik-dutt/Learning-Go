package main

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("Hello!")
	kartik := &Person{
		FirstName: "Kartik",
		LastName:  "Dutt",
		Age:       18,
	}

	data, _ := proto.Marshal(kartik)
	fmt.Println(data)

	newKartik := &Person{}
	_ = proto.Unmarshal(data, newKartik)
	fmt.Println(newKartik)
}
