package main

import "fmt"

func sum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}


func main(){
	var nums = []int {1,2,3,4,5}
	fmt.Println(sum(nums))
}
