package main

import (
	"fmt"
	"time"
)

func main() {
	print := fmt.Println
	then := time.Date(2020, 12, 31, 11, 59, 59, 0, time.UTC)
	print(then)
	print(then.Location())
	print(then.After(time.Now()))
	print(time.Now().Sub(then))
	print(then.Add(time.Now().Sub(then)))
}
