package a03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestA03(t *testing.T) {
	res := A03("bingo")
	exp := "your string is bingo"
	assert.Equal(t, exp, res, "test01 ")
	//assert.NoError(t, false, "There are no errors")

	res = A03("bingo")
	exp = "your string is bingo"
	assert.Equal(t, exp, res, "test01 ")

}
