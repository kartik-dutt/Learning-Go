package main

import "fmt"

func main() {
	ht := make(map[string]int)
	ht["Kartik"] = 1
	fmt.Println(ht["Kartik"])
	fmt.Println(len(ht))
	delete(ht, "Kartik")
	_, present := ht["Kartik"]
	if present {
		fmt.Println("Present")
	} else {
		fmt.Println("Not present")
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
}
