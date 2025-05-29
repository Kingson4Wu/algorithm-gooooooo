package string_match

import (
	"fmt"
	"testing"
)

func TestMatch(t *testing.T) {

	A := "abcdebcdf"
	B := "bcd"

	fmt.Println("BF查找结果:", BFSearch(A, B)) // 输出 1
	fmt.Println("RK查找结果:", RKSearch(A, B)) // 输出 1

}
