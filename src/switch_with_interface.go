package main

import "fmt"

func test_switch_fun(s interface {}) string {

	switch s.(type) {
	case string:
		return "string"
	case int:
		return "int"
	case bool:
		return "bool"
	default:
		return "default"
	}
}


func main(){
	var a interface{} // a pointer, rather than a generic type
	a = 4
	fmt.Println(test_switch_fun(a))
	a = "hello"
	fmt.Println(test_switch_fun(a))
}
