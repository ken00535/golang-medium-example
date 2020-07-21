package postgres

import (
	"encoding/json"
	"io/ioutil"
)

type dbConfig struct {
	Addr     string
	Port     int
	Username string
	Name     string
	Password string
}

func getDbConfig() *dbConfig {
	config := dbConfig{}
	file := "./configs/postgres/config.json"
	data, err := ioutil.ReadFile(file)
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	return &config
}
