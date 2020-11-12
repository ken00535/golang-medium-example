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

type UserRepository interface {
	// 根據使用者id查詢得到一個使用者或是錯誤資訊
	FindOne(id int) (User, error)
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
