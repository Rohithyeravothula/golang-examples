package main

import "fmt"

func fibonacci(n int) int {
	if n<1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}


func main(){
	fmt.Println(fibonacci(5))
	// since this is a recursion based solution to fibonacci, don't try to
	// run for large numbers as it might take up all you system's resources
}
