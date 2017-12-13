package lang

import (
	list "container/list"
	"fmt"
	"time"
)

var ll *list.List

func init() {
	ll = new(list.List)
	ll.Init()
	ll.PushBack(1)
	ll.PushBack(2)
	ll.PushBack(3)
	ll.PushBack(4)
	ll.PushBack(5)
	ll.PushBack(6)
}

func ListBack(){
	for {
		e := ll.Back()
		fmt.Println(e)
		ll.Remove(e)

		time.Sleep(1 * time.Second)
	}

}
