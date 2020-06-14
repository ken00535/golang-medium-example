package main

import (
	"fmt"
	"router/pkg/route"
)

func main() {
	router := route.NewRouter()
	router.Add("hello", helloHandler)
	var res route.Message
	req := route.Message{
		Identification: "hello",
		Content:        "Gopher",
	}
	router.Run(&res, &req)
	fmt.Println(res.Content)
}

func helloHandler(res, req *route.Message) {
	res.Content = req.Content
}
