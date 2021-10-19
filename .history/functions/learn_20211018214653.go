package functions

import "fmt"

var Fullname string


func add(x int, y int) {
	return x + y
}

func Learn() {
	fmt.Println(add(10, 5))
}

