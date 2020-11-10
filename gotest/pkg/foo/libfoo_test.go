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
