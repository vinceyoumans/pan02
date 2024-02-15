package a01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestT02(t *testing.T) {
	res := A01("bingo")
	exp := "your string is bingo"
	assert.Equal(t, exp, res, "test01 ")
	//assert.NoError(t, false, "There are no errors")

	res = A01("bingo")
	exp = "your string is bingo"
	assert.Equal(t, exp, res, "test01 ")

}
