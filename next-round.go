package main 

import "fmt"

func main(){
	var(
		n, k int
	)
	
	fmt.Scan(&n, &k)
	
	scores := make([]int,n)
	
	for i:=0; i<n; i++{
		fmt.Scan(&scores[i])
	}
	
	thresold := scores[k-1]
	
	count := 0
	
	for x:=0; x<n; x++{
		if scores[x] >= thresold && scores[x]>0{
			count++
		}
	}
	
	fmt.Println(count)
	
}