package main

import "fmt"

func test_switch(s string) bool {
	switch s {
	case "hello":
		return true
	case "hello world":
		return false
	default:
		return false
	}
}


func main() {
	fmt.Println(test_switch("hello"))
}