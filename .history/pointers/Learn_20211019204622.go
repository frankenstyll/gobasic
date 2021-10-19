package pointers

import "fmt"

func Learn() {

	//	ตัวแปร ที่เก็บ memory address ของตัวแปรอื่น
	x := 10
	fmt.Println(x)

	fmt.Println(&x) //0xc00000c098 , & คือจะได้ memory address

	y := x //จะได้ค่า x =10
	fmt.Println(y)
	fmt.Println(&y) //0xc00000c0e0 , จะได้ memory address ใหม่

	fmt.Println("-------")

	//var p *int = &x
	p := &x        //ชี้ไปหา pointer x
	fmt.Println(p) //0xc00000c098 , จะได้ memory address ของ x

	fmt.Println(*p) // * จะได้ค่า 10 ซึ่งเป็นค่าของ x , อ่านค่า x ผ่าน p pointer , ใน go เรียกว่า dereference

	*p = 20 //assign ค่าใหม่
	fmt.Println(x)
}
