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
	var schemas []schema
	data, err := ioutil.ReadFile("../../config/callback.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &schemas)
	if err != nil {
		panic(err)
	}
	t := template.Must(template.ParseFiles("../../tmpl/main.tmpl", "../../tmpl/callbackTemplate.tmpl", "../../tmpl/contextTemplate.tmpl"))
	t.Execute(os.Stdout, schemas)
}
