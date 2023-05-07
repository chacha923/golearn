package skiplist

import (
	"testing"

	"math/rand"
)

func TestSL(t *testing.T) {
	sl := NewSkipList()

	for i := 0 ; i < 30 ; i++  {
		sl.Add(string(i),"foo")
	}
	sl.Print()
	println("\n\n----------\n\n\n")
	if sl.Find("23") != nil{
		println("\nok")
	}else{
		println("\nfalse")
	}

	sl.Del("0")
	sl.Print()
	
}

func TestRand(t *testing.T){

	println(rand.Float32())
}