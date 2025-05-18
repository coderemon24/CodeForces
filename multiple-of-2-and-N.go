package main

import "fmt"

func multipleOf2AndN(N int) int {
	for i := N; ; i += 1 {
		
		if i % 2 == 0 && i % N == 0 {
			return i
		}
	}
}

func main() {
	var N int
	fmt.Scan(&N)

	fmt.Print(multipleOf2AndN(N))
}