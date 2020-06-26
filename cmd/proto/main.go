package main

import (
	"example/pkg/pb/person"

	"github.com/golang/protobuf/proto"
)

func main() {
	data := person.Person{}
	err := proto.Unmarshal([]byte(*msg), &data)
}
