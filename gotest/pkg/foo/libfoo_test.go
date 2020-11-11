package foo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFooBasic(t *testing.T) {
	want := 2
	actual := fooBasic(1, 1)
	assert.Equal(t, want, actual)
}

func TestFooDatabase(t *testing.T) {
	want := 1
	getUserAge = func() int {
		return 1
	}
	actual := fooDatabase()
	assert.Equal(t, want, actual)
}
