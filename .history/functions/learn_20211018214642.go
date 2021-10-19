package functions

import "fmt"

var Fullname string

func Learn() {
	fmt.Println(add(10, 5))
}

func add(x int, y int) {
	return x + y
}
