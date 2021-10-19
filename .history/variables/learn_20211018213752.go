package variables

import (
	"fmt"
	"strconv"
)

var fullname string // เป็น global
var email = "nontapap.th@gmail.com"
var c, python bool = true, false

//constant จะประกาศแบบ := ไม่ได้นะ
const vat int = 7 //ตัวแปรชื่อ vat เป็น int มีค่า 7

func Learn() {

	fullname = "Franken styll"

	// fmt.Println("Learn var")
	fmt.Println(fullname)
	fmt.Println(email)

	fmt.Printf("Hello %s Email %s \n", fullname, email)
	fmt.Println(c, python)

	//ประกาศตัวแปรแบบสั้นๆ ต้องประกาศใน function
	companyName := "KTB"
	isShow := true
	age := 10
	fmt.Println(companyName, isShow, age)

	//เช็คว่าตัวแปรเป็น type อะไร
	fmt.Printf("isShow is %T \n", isShow)

	//แปลง int ไป float64
	f1 := float64(age)
	fmt.Printf("f1 is %T \n", f1)

	//แปล float เป็น string
	s1 := strconv.i
}
