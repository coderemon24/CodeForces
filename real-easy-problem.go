package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scanf("%d", &n)
		fmt.Println(n - 1)
	}
}
