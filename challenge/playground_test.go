package challenge

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	assert.Equal(t, int64(123), reverse(321))
	assert.Equal(t, int64(43), reverse(34))
	assert.Equal(t, int64(3), reverse(3))
}

func TestBinGap(t *testing.T) {
	assert.Equal(t, int(2), binGap(9))
	assert.Equal(t, int(4), binGap(529))
	assert.Equal(t, int(0), binGap(15))
	assert.Equal(t, int(0), binGap(32))
	assert.Equal(t, int(0), binGap(1))
}

func TestFindTree(t *testing.T) {
	assert.Equal(t, int64(2), findTree(Relation{0, 1}, Relation{0, 2}, Relation{3, 4}))
	assert.Equal(t, int64(1), findTree(Relation{0, 4}, Relation{0, 3}, Relation{0, 4}))
	assert.Equal(t, int64(1), findTree(Relation{0, 4}, Relation{0, 3}, Relation{0, 1}))
	assert.Equal(t, int64(3), findTree(Relation{0, 1}, Relation{0, 2}, Relation{3, 4}, Relation{4, 4}))
}

func TestPairOfSum(t *testing.T) {
	assert.Equal(t, "first pair is 4 and 1", pairOfSum([]int64{0, 3, 4, 1, 3}, 5))
	assert.Equal(t, "first pair is 0 and 3", pairOfSum([]int64{0, 3, 4, 1, 3}, 3))
	assert.Equal(t, "first pair is 1 and 3", pairOfSum([]int64{0, 3, 4, 1, 3}, 4))
	assert.Equal(t, "first pair is 3 and 1", pairOfSum([]int64{3, 1, 4, 1, 3}, 4))
	assert.Equal(t, "not found", pairOfSum([]int64{3, 1, 4, 1, 3}, 6))
}
