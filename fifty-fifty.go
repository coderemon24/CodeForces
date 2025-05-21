package main

import (
	"fmt"
)

func checkString(s string) string {
	if len(s) != 4 {
		return "No"
	}
	
	count := make(map[rune]int)
	for _, char := range s {
		count[char]++
	}
	
	if len(count) != 2 {
		return "No"
	}
	
	for _, v := range count {
		if v != 2 {
			return "No"
		}
	}
	
	return "Yes"
	
}


func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(checkString(s))
}