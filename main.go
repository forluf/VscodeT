package main

import "fmt"

type men struct {
	name string
	age  int
}

func main() {
	var a men
	a = men{name: "Aex", age: 18}
	fmt.Printf("%+v", a)
}
