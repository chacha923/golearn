package lang

import (
	"fmt"
	"log"
)

func RunError() int {
	defer func() {
		fmt.Println("first defer")
		log.Println("method will run return first defer")
	}()

	defer func() {
		fmt.Println("second defer")
		log.Println("method will run return second defer")
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in defer before a panic")
		}
		log.Println("method will run return third defer before a panic")
	}()

	//panic("has panic")
	fmt.Println("a line after panic")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in forth defer after a panic")
		}
		log.Println("method will run return from forth defer after a panic")
	}()

	log.Println("it will run return in last line and return 1")
	return 1

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in fifth defer after last return")
		}
		log.Println("method will run return fifth defer before a panic")
	}()
	return 2
}
