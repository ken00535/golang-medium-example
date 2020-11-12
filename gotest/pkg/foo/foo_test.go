package foo

import (
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/jinzhu/gorm"
	"github.com/prashantv/gostub"
	"github.com/stretchr/testify/assert"
)

func TestFooBasic(t *testing.T) {
	expect := 2
	actual := fooBasic(1, 1)
	assert.Equal(t, expect, actual)
}

func TestFooDatabase1(t *testing.T) {
	want := 1
	stub := gostub.Stub(&getUserAge1, func() int {
		return 1
	})
	defer stub.Reset()
	actual := fooDatabaseCase1()
	assert.Equal(t, want, actual)
}

func TestFooDatabase2(t *testing.T) {
	want := 1
	patch := monkey.Patch(getUserAge2, func() int {
		return 1
	})
	defer patch.Restore()
	actual := fooDatabaseCase2()
	assert.Equal(t, want, actual)
}

func TestFooDatabase3(t *testing.T) {
	want := 1
	user := User{Age: 1}
	db := &gorm.DB{}
	patch := monkey.Patch(gorm.Open, func(string, ...interface{}) (*gorm.DB, error) {
		return db, nil
	})
	patchFirst := monkey.PatchInstanceMethod(reflect.TypeOf(db), "First", func(_ *gorm.DB, out interface{}, _ ...interface{}) *gorm.DB {
		val := reflect.ValueOf(out).Elem()
		substitute := reflect.ValueOf(user)
		val.Set(substitute)
		return db
	})
	patchClose := monkey.PatchInstanceMethod(reflect.TypeOf(db), "Close", func(*gorm.DB) error {
		return nil
	})
	defer func() {
		patch.Restore()
		patchFirst.Restore()
		patchClose.Restore()
	}()
	actual := fooDatabaseCase3()
	assert.Equal(t, want, actual)
}

func TestFooDatabase4(t *testing.T) {
	want := 1
	db := newDbMock()
	actual := fooDatabaseCase4(db)
	assert.Equal(t, want, actual)
}
