package variables

import "fmt"

var fullname string // เป็น global

func Learn() {

	fullname = 'asd'

	fmt.Println("Learn var")
	fmt.Println(fullname)

}
