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
