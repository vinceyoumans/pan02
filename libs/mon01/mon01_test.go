package mon01

import (
	"github.com/stretchr/testify/assert"
	"github.com/vinceyoumans/pan02/structs"
	"testing"
)

func TestMon01(t *testing.T) {

	m := structs.Mon{}
	m.DirToMonitor = "./temp"
	m.DirTorecordJSON = "./json"
	m.ConcMan = 3

	res := Mon01(m)
	exp := "your string is bingo"
	assert.Equal(t, exp, res, "test01 ")
	//assert.NoError(t, false, "There are no errors")

	//res = Mon01(m)
	//exp = "your string is bingo"
	//assert.Equal(t, exp, res, "test01 ")

}
