package main

import "fmt"

func main() {
	var N, T, A int 
	fmt.Scan(&N, &T, &A)

	remaining := N - (T + A)

	if T > A + remaining || A > T + remaining {
		fmt.Println("Yes") 
	} else {
		fmt.Println("No")
	}
}
