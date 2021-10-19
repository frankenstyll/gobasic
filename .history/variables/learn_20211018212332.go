package variables

import "fmt"

var fullname string // เป็น global
var email = ""
func Learn() {

	fullname = "Franken styll"

	fmt.Println("Learn var")
	fmt.Println(fullname)

}
