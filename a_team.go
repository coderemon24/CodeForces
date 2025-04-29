package main
import(
	"fmt"
)

func main(){
	var n int
	fmt.Scan(&n)
	var count = 0
	for i:= 0; i<n; i++{
		var x,y,z int
		fmt.Scan(&x,&y,&z)
		if x+y+z >= 2{
			count++
		}
	}
	fmt.Println(count)
}