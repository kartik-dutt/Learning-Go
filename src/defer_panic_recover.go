package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func panicker() {
	// Panic is same as exception.
	div := func(a, b float32) (float32, error) {
		if b != 0 {
			return a / b, nil
		}

		return -1, errors.New("Division by 0 isn't possible")
	}

	defer func() {
		if ez := recover(); ez != nil {
			log.Println("Error :", ez)
		}
	}()

	_, erro := div(3, 0.0)
	if erro != nil {
		// All defered tasks execute here. because of panic is return statement.
		panic(erro.Error())
	}

	// This won't be printed since the recover function will make the current function exit.
	fmt.Println("Hello World")
}

func main() {
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
	// It looks it for any defered function.
	// LIFO order is followed for defered.
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	// Defer used to close resources.
	res, err := http.Get("http://www.google.com/robots.txt")
	// We can access res below but we ensure it's closing.
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

	// Fun with defer.
	a := "start"
	// Defer takes paramter when it reads it.
	defer fmt.Println(a)
	a = "end"

	panicker()
}
