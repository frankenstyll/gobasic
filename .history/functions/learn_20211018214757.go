package functions

import "fmt"

var Fullname string

//หรือ func add(x, y int) ถ้า x, y ม
func add(x int, y int) int {
	return x + y
}

func Learn() {
	fmt.Println(add(10, 5))
}
