package main

import (
	"fmt"
)

func main() {
	var c rune
	fmt.Scanf("%c", &c)  

	switch c {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Println("vowel") 
	default:
		fmt.Println("consonant")
	}
}
