package main

import "example/pkg/postgres"

func main() {
	client := postgres.DBClient{}
	client.Connect()
	client.Insert()
}
