package dp

import "testing"

func TestStair(t *testing.T) {
	if stair(5) != 8 {
		t.Error("stair error")
	}
}
