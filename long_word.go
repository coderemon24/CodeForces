package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	var n int 
	fmt.Scan(&n)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i <= n; i++ {
		scanner.Scan()
		word := strings.TrimSpace(scanner.Text())
		if len(word) > 10 {
			abbr := fmt.Sprintf("%c%d%c", word[0], len(word)-2, word[len(word)-1])
			fmt.Println(abbr)
		}else{
			fmt.Println(word)
		}
	}
}