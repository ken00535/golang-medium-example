package main

import "fmt"

//go:generate bash -c "go run ../codegen/main.go > ./context.go"

func main() {
	printNum := func(num int) {
		fmt.Println(num)
	}
	context := &Context{}
	context.OnClick(printNum)
	context.EmitClick(5)
}
