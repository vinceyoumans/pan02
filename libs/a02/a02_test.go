package a02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestA02(t *testing.T) {
	res := A02("bingo")
	exp := "your string is bingo"
	assert.Equal(t, exp, res, "test A02 ")
	//assert.NoError(t, false, "There are no errors")

	res = A02("bingo")
	exp = "your string is bingo"
	assert.Equal(t, exp, res, "test A02 ")

}
