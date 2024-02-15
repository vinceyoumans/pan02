package f01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestF01(t *testing.T) {
	res := F01("bingo")
	exp := "your string is bingo"
	assert.Equal(t, exp, res, "test01 ")
	//assert.NoError(t, false, "There are no errors")

	res = F01("bingo")
	exp = "your string is bingo"
	assert.Equal(t, exp, res, "test01 ")

}
