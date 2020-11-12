package foo

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

// Database is database
type Database interface {
	First(interface{})
}

func fooBasic(num1 int, num2 int) int {
	return num1 + num2
}

func fooDatabaseCase1() int {
	return getUserAge1()
}

func fooDatabaseCase2() int {
	return getUserAge2()
}

func fooDatabaseCase3() int {
	var user User
	db, err := gorm.Open("postgres", "host=myhost user=gorm dbname=gorm sslmode=disable password=mypassword")
	if err != nil {
		panic("connect fail")
	}
	res := db.First(&user)
	if res.Error != nil {
		panic("error")
	}
	db.Close()
	return user.Age
}

func fooDatabaseCase4(db Database) int {
	var user User
	db.First(&user)
	return user.Age
}

var getUserAge1 = func() int {
	var user User
	db, err := gorm.Open("postgres", "host=myhost user=gorm dbname=gorm sslmode=disable password=mypassword")
	if err != nil {
		panic("connect fail")
	}
	res := db.First(&user)
	if res.Error != nil {
		panic("error")
	}
	db.Close()
	return user.Age
}

func getUserAge2() int {
	var user User
	db, err := gorm.Open("postgres", "host=myhost user=gorm dbname=gorm sslmode=disable password=mypassword")
	if err != nil {
		panic("connect fail")
	}
	res := db.First(&user)
	if res.Error != nil {
		panic("error")
	}
	db.Close()
	return user.Age
}
