package main

import (
	"example/pkg/postgres"
	"fmt"
)

func main() {
	client := postgres.DBClient{}
	client.Connect()
	player := postgres.Player{
		Age:      18,
		Username: "ken",
		Budget:   1000,
	}
	client.Insert(player)
	players, err := client.Get()
	if err != nil {
		fmt.Println(err)
	}
	if len(players) > 1 {
		players[1].Budget = 2000
		client.Update(players[1])
	}
	client.Disconnect()
}
