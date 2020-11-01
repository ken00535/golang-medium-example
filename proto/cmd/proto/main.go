package main

import (
	"encoding/json"
	"example/pkg/person"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
)

func main() {
	// write file
	msgWrite := person.Person{}
	msgWrite.Age = 22
	msgWrite.Gender = "Female"
	msgWrite.Name = "Cythia"
	msgWrite.Address = "Boston ,US"
	dataWrite, err := proto.Marshal(&msgWrite)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("encode.txt", dataWrite, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	// write json
	dataJSON, err := json.Marshal(&msgWrite)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("encode_json.txt", dataJSON, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	// read file
	data, _ := ioutil.ReadFile("encode.txt")
	msg := person.Person{}
	err = proto.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(msg.Name)
	fmt.Println(msg.Gender)
	fmt.Println(msg.Address)
	fmt.Println(msg.Age)
}
