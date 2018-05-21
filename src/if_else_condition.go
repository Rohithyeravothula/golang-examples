package main

import "fmt"

func main(){
	s := "hello" // another way to declare a variable, it asserts type from the initialization value
	if s == "hello world" {
		fmt.Println("hello world")
	} else if s == "hello" {
		fmt.Println("missing world")
	} else {
		fmt.Println("missing everything")
	}
}
