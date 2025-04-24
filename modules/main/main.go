package main

import (
	"fmt"
	"gomaster/modules/somemodule/somepackage"
	"gomaster/modules/ypmodule/calc"
)

func main() {
	fmt.Println(calc.AddInts(1, 2))
	somepackage.SomeFunc()
}
