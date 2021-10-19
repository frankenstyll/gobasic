package variables

import "fmt"

var fullname string // เป็น global
var email = "nontapap.th@gmail.com"
var c, python bool = true, false

func Learn() {

	fullname = "Franken styll"

	fmt.Println("Learn var")
	fmt.Println(fullname)
	fmt.Println(email)

	fmt.Printf("Hello %s Email %s \n", fullname, email)
	fmt.Println(c, python)
}
