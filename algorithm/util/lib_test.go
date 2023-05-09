package util

import (
	"fmt"
	"testing"
)

func TestDefaultValue(t *testing.T) {
	fmt.Println(DefaultValue[int64](1))
}
