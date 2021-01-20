package main

import (
	"fmt"
	"time"
)

func main() {
	print := fmt.Println
	print(time.Now().Unix())
	print(time.Now().UnixNano())
}
