package main

import (
	"fmt"
)

func main() {
	var X, t int
	fmt.Scan(&X, &t)

	remaining := X - t
	if remaining < 0 {
		remaining = 0
	}
	fmt.Println(remaining)
}
