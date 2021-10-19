package functions

import "fmt"

var Fullname string

func Learn() {
	fmt.Println(add())
}

func add(x int, y int) {
	return x + y
}
