package main

import "fmt"

func maxDays(n int, m int) int {
	days := 0
	socks := n
	
	for socks > 0 {
		days++
		socks--
		
		if days % m == 0 {
			socks++
		}
	}
	
	return days
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Print(maxDays(n, m))
}