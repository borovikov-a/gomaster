package remotepkg

import (
	"fmt"
)

func AddAndPrint(a, b int) int {
	res := a + b
	fmt.Println("This function just have ")
	return res
}
