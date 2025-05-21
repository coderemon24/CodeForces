package main

import "fmt"

func distributeCrackers(N int, K int) int64 {
	if K == 0 {
		return 0
	}
	
	if N % K == 0 {
		return 0
	}
	
	return 1
}

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	fmt.Print(distributeCrackers(N, K))
}