package lang

import (
	"math"
	"fmt"
)

func Int64ToInt32(){
	a := int(math.MaxInt32 + 1)
	var b int32
	b = int32(a)
	fmt.Println(b)
}
