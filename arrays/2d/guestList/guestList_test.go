package guestList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	guestList := [10][2]string{}
	guestList = Generate()
	assert.Equal(t, 10, len(guestList))
}
