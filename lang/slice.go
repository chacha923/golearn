package lang

import "fmt"

func Slice(){
	slice := make([]int,0,10)
	fmt.Println(len(slice))
	slice = append(slice,1)
	fmt.Println(slice)
}
