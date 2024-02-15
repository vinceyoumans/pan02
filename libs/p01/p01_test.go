package p01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestP01(t *testing.T) {
	res := P01("bingo")
	exp := "your string is bingo"
	assert.Equal(t, exp, res, "test01 ")
	//assert.NoError(t, false, "There are no errors")

	res = P01("bingo")
	exp = "your string is bingo"
	assert.Equal(t, exp, res, "test01 ")

}
