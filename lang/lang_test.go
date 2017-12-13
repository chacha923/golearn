package lang

import (
	"testing"
	"runtime"
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags|log.Lshortfile)
}

func TestRangeMap(t *testing.T) {
	RangeMap()
}

func TestRangeArrayPoint(t *testing.T) {
	RangeArrayPoint()
}

func TestRangeArray(t *testing.T) {
	RangeArray()
}

func TestMaxProc(t *testing.T) {
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
}

func TestListBack(t *testing.T) {
	ListBack()
}

func TestRunError(t *testing.T) {
	fmt.Println(RunError())
}
