package main

import "fmt"

func soldierAndBananas(k, n, w int) int {
	totalCost := 0

	for i := 1; i <= w; i++ {
		totalCost += i * k
	}

	borrow := totalCost - n
	if borrow < 0 {
		borrow = 0
	}

	return borrow
}

func main() {
	var k, n, w int
	fmt.Scan(&k, &n, &w)

	fmt.Print(soldierAndBananas(k, n, w))
}