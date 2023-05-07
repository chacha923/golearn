package test

import "testing"

func TestStrCompare(t *testing.T) {
	a := "21"
	b := "23"
	if a <= b {
		println("a <= b")
		return
	}
	println("a > b")
	return

}
