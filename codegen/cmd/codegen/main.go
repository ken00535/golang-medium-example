package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

	funcLowerCase := template.FuncMap{"lower": strings.ToLower}
	t := template.Must(template.New("main.tmpl").Funcs(funcLowerCase).ParseFiles("../../tmpl/main.tmpl"))
	t = template.Must(t.ParseFiles("../../tmpl/callbackTemplate.tmpl"))
	t = template.Must(t.ParseFiles("../../tmpl/contextTemplate.tmpl"))
	err = t.Execute(os.Stdout, schemas)
	if err != nil {
		fmt.Println(err)
	}
}
