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

func (m *DBClient) Insert(player Player) error {
	res := m.client.Create(&player)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *DBClient) Get() ([]Player, error) {
	players := []Player{}
	res := m.client.Order("id").Find(&players)
	if res.Error != nil {
		return nil, res.Error
	}
	return players, nil
}

func (m *DBClient) Update(player Player) error {
	res := m.client.Save(&player)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *DBClient) Delete(player Player) error {
	res := m.client.Delete(&player)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
