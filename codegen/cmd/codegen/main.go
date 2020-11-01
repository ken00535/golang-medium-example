package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"text/template"
)

type schema struct {
	EventName   string
	CallbackArg string
}

func main() {
	var schema schema
	data, err := ioutil.ReadFile("./config/callback.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &schema)
	if err != nil {
		panic(err)
	}
	t := template.Must(template.ParseFiles("./tmpl/main.tmpl"))
	t.Execute(os.Stdout, schema)
}
