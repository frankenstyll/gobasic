package variables

import "fmt"

var fullname string // เป็น global

func Learn() {

	fullname = "Franken stylle"

	fmt.Println("Learn var")
	fmt.Println(fullname)

}
