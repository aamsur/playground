package playground

import (
	"testing"
	"fmt"
)

func TestReverse(t *testing.T) {
	fmt.Println("reverse result :")
	fmt.Println(reverse(321))
	fmt.Println(reverse(34))
	fmt.Println(reverse(3))
}

func TestBinGap(t *testing.T) {
	fmt.Println("Binary Gap result :")
	fmt.Println(binGap(9))
	fmt.Println(binGap(529))
	fmt.Println(binGap(15))
	fmt.Println(binGap(32))
	fmt.Println(binGap(1))
}

func TestFindTree(t *testing.T) {
	fmt.Println("Find tree result :")

	fmt.Println(findTree(Relation{0, 1}, Relation{0, 2}, Relation{3, 4}))
	fmt.Println(findTree(Relation{0, 4}, Relation{0, 3}, Relation{0, 4}))
	fmt.Println(findTree(Relation{0, 4}, Relation{0, 3}, Relation{0, 1}))
	fmt.Println(findTree(Relation{0, 1}, Relation{0, 2}, Relation{3, 4}, Relation{4, 4}))
}

func TestPairOfSum(t *testing.T) {
	fmt.Println("Find pair of sum result :")

	fmt.Println(pairOfSum([]int64{0, 3, 4, 1, 3}, 5))
	fmt.Println(pairOfSum([]int64{0, 3, 4, 1, 3}, 3))
	fmt.Println(pairOfSum([]int64{0, 3, 4, 1, 3}, 4))
	fmt.Println(pairOfSum([]int64{3, 1, 4, 1, 3}, 4))
}
