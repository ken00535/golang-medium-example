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

func fooDatabaseCaseByValueFunc() int {
	return getUserAgeValueFunc()
}

func fooDatabaseCaseByFunc() int {
	return getUserAgeFunc()
}

func fooDatabaseCaseDirectCall() int {
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

func fooDatabaseCaseIndirectCall(db Database) int {
	var user User
	db.First(&user)
	return user.Age
}

var getUserAgeValueFunc = func() int {
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

func getUserAgeFunc() int {
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
