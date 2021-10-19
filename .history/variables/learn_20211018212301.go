package variables

import "fmt"

var fullname string // เป็น global

func Learn() {

	fullname = ""

	fmt.Println("Learn var")
	fmt.Println(fullname)

}
