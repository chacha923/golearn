package lang

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"testing"
	"time"
)

var array [3]int

func init() {
	array[0] = 0
	array[1] = 1
	array[2] = 2
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestRangeMap(t *testing.T) {
	RangeMap()
}

func TestRangeSlice(t *testing.T) {
	slice := NewSlice()
	fmt.Println(RangeSlice(slice))
}

func TestRangeArrayPoint(t *testing.T) {
	fmt.Println(RangeArrayPoint(array))
}

func TestRangeArray(t *testing.T) {
	fmt.Println(RangeArray(array))
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

func TestSlice(t *testing.T) {
	Slice()
}

func TestTmp(t *testing.T) {
	fmt.Println("temp dir : ", os.TempDir())
}

func TestBar(t *testing.T) {
	bar()
}

func TestCh(t *testing.T) {
	var mq = NewMQ()
	go mq.Run()
	time.Sleep(1 * time.Second)
	mq.Close()
}
