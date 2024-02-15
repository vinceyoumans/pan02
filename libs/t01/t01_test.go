package t01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestT01(t *testing.T) {
	res := T01("bingo")
	exp := "your string is bingo"
	assert.Equal(t, exp, res, "test01 ")
	//assert.NoError(t, false, "There are no errors")

	res = T01("bingo")
	exp = "your string is bingo"
	assert.Equal(t, exp, res, "test01 ")

}
