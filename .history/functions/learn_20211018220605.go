package functions

import (
	"fmt"
	"strings"
)

var Fullname string

func Learn() {
	fmt.Println(add(10, 5))
	fmt.Println(converter("go lang", "non@gmail"))

	sum(10, 20, 30, 40)
}

//หรือ func add(x, y int) ถ้า x, y มีค่าเท่ากันก็ประกาศแบบนี้ได้
func add(x int, y int) int {
	//return int
	return x + y
}

func converter(title, email string) (string, string) {
	//return 2 ค่า
	title = strings.ToUpper(title)
	em := strings.ToUpper(email)
	return title, em
}

///...int เรียกว่า Variadic function
func sum(numbers ...int) {
	// numbers จะถูกแปลงเป็น  array
	fmt.Println(numbers)

	total := 0
	// go จะมีแค่ loop for
	for _, num := range numbers {

		total += num
	}
	fmt.Println(total)
}

func sum/(numbers ...int) {
	// numbers จะถูกแปลงเป็น  array
	fmt.Println(numbers)

	total := 0
	// go จะมีแค่ loop for
	for _, num := range numbers {

		total += num
	}
	fmt.Println(total)
}
