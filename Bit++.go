package main
import "fmt"
func main(){
	var n int
	fmt.Scan(&n)
	var x = 0;
	for i := 0; i < n; i++ {
		var input string
		fmt.Scan(&input)
		if input == "++x" || input == "x++" {
			x += 1
		}else if input == "--x" || input == "x--" {
			x -= 1
		}
	}
	fmt.Println(x)
}