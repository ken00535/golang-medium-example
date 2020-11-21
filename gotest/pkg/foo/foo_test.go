package foo

import (
	"reflect"
	"testing"

	"bou.ke/monkey"
	gomock "github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/prashantv/gostub"
	"github.com/stretchr/testify/assert"
)

func TestFooBasic(t *testing.T) {
	expect := 2
	actual := fooBasic(1, 1)
	assert.Equal(t, expect, actual)
}

func TestFooDatabaseByValueFunc(t *testing.T) {
	want := 1
	stub := gostub.Stub(&getUserAgeValueFunc, func() int {
		return 1
	})
	defer stub.Reset()
	actual := fooDatabaseCaseByValueFunc()
	assert.Equal(t, want, actual)
}

func TestFooDatabaseByFunc(t *testing.T) {
	want := 1
	patch := monkey.Patch(getUserAgeFunc, func() int {
		return 1
	})
	defer patch.Restore()
	actual := fooDatabaseCaseByFunc()
	assert.Equal(t, want, actual)
}

func TestFooDatabaseByMonkeyPatch(t *testing.T) {
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
	actual := fooDatabaseCaseDirectCall()
	assert.Equal(t, want, actual)
}

func TestFooDatabaseCustomMock(t *testing.T) {
	want := 1
	m := newDbMock()
	actual := fooDatabaseCaseIndirectCall(m)
	assert.Equal(t, want, actual)
}

func TestFooDatabaseGomock(t *testing.T) {
	want := 1
	var user User
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockDatabase(ctrl)
	m.EXPECT().First(gomock.Eq(&user)).SetArg(0, User{Age: 1})
	actual := fooDatabaseCaseIndirectCall(m)
	assert.Equal(t, want, actual)
}
