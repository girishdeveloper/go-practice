package demo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const newValue = "working here"

func UpdateArray(newArray *[2]string) {
	newArray[1] = newValue //+ "123"
}

func TestUpdateArray(t *testing.T) {
	testArray := [2]string{"Value1", "Value2"}
	UpdateArray(&testArray)
	assert.Equal(t, newValue, testArray[1])
}
