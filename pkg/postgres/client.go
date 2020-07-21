package postgres

import (
    "github.com/jinzhu/gorm"
    // for postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
)

type DBClient struct {
	client *gorm.DB
}

func (m *DBClient) Connect() {
	config := getDbConfig()
	client, err := gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			config.Addr,
			config.Port,
			config.Username,
			config.Name,
			config.Password,
		),
	)
	if err != nil {
		panic(err)
	}
	m.client = client
}

func (m *DBClient) Disconnect() {
	m.client.Close()
}

func (m *DBClient) Insert() {
	player := Player{}
	m.client.Create(&player)
}