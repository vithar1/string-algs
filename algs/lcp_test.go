package algs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinBlockSplit(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 1, 0}
	blocks := minBlockSplit(arr, 3)
	assert.Equal(t, []int{0, 3, 7}, blocks)
}

func TestSparceTable(t *testing.T) {
	arr := []int{2, 1, 0, 3}
	sp := sparseTable(arr)
	assert.Equal(t, [][]int{
		{0, 1, 2},
		{1, 2, 2},
		{2, 2, 2},
		{3, 3, 3},
	}, sp)
}

func TestMinFromST(t *testing.T) {
	arr := []int{2, 1, 0, 3}
	sp := sparseTable(arr)
	assert.Equal(t, 0, minFromST(arr, sp, 0, 0))
	assert.Equal(t, 1, minFromST(arr, sp, 0, 1))
	assert.Equal(t, 2, minFromST(arr, sp, 0, 2))
	assert.Equal(t, 2, minFromST(arr, sp, 1, 2))
	assert.Equal(t, 2, minFromST(arr, sp, 2, 2))
	assert.Equal(t, 2, minFromST(arr, sp, 0, 3))
	assert.Equal(t, 3, minFromST(arr, sp, 3, 3))
	assert.Equal(t, 2, minFromST(arr, sp, 1, 3))
}

func TestBTNext(t *testing.T) {
	bt := []bool{false, false, false}
	btNext(bt)
	assert.Equal(t, []bool{false, false, true}, bt)
	btNext(bt)
	assert.Equal(t, []bool{false, true, false}, bt)
	btNext(bt)
	assert.Equal(t, []bool{false, true, true}, bt)
	btNext(bt)
	assert.Equal(t, []bool{true, false, false}, bt)
	btNext(bt)
	assert.Equal(t, []bool{true, false, true}, bt)
}

func TestBlockMinTable(t *testing.T) {
	assert.Equal(t, [][][]int{
		{{0, 1, 2}, {0, 1, 2}, {0, 0, 2}},
		{{0, 1, 1}, {0, 1, 1}, {0, 0, 2}},
		{{0, 0, 0}, {0, 1, 2}, {0, 0, 2}},
		{{0, 0, 0}, {0, 1, 1}, {0, 0, 2}},
	}, blockMinTable(2))
}
