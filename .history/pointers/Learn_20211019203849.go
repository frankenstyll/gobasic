package pointers

import "fmt"

func Learn() {

	//	ตัวแปร ที่เก็บ memory address ของตัวแปรอื่น
	x := 10
	fmt.Println(x)

	fmt.Println(&x) //0xc00000c098 , & คือจะได้ memory address

	y := x //จะได้ค่า x =10
	fmt.Println(y)
}
