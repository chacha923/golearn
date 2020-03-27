package c2go

/*
#include <stdio.h>

void printint(int v) {
    printf("printint: %d\n", v);
}
*/
import "C"

func ExampleCGO() {
	v := 42
	C.printint(C.int(v))
}
