package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type men struct {
	Name string
	Age  int
}

func main() {
	var a men
	a = men{Name: "Alex", Age: 18}
	bytes, err := json.Marshal(a)
	fmt.Printf("%v", bytes)
	errorOnExit(err)
}

func errorOnExit(err error) {
	if err == nil {
		fmt.Println()
		fmt.Println("No errors")
	} else {
		fmt.Println(err)
		os.Exit(1)
	}
}
