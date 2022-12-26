package main

import "fmt"

func main() {
	r := sum(4, 6, 2)
    fmt.Printf("sum is : %d \n", r)
}

func sum(integers ...int64) int64 {
	var t int64
	for _, i := range integers {
		t += i
	}

	return t
}
