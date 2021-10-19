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

	*p = 20 //assign ค่าใหม่  ค่าจะถูก assign ไปยัง x
	fmt.Println(x)

	fmt.Println("-------")
	//การเล่นกับ pointer ด้วย struct
	b := student{name: "Bob"}
	fmt.Println(b)

	setName(b)
	fmt.Println(b) //ค่ายังเป็น Bob อยู่ ไม่เปลี่ยนเป็น Boy

	fmt.Println("************************************")
	setName2(&b) //ส่ง pointer address เข้าไปให้ setName2 แล้วให้ setName2 รับเป็น struct ค่า ถึงจะเปลี่ยนค่าได้
	fmt.Println(b)
	fmt.Println("************************************")
	/*
		สรุป
		& เป็น operator ถ้าวางไว้หน้าตัวแปร เราจะได้ memory address
		* ถ้าวางไว้หน้าตัวแปร เป็นการดึงค่า value ที่ memory address นั้นๆ
		* ถ้าวางไว้หน้า Type เช่น struct 
	*/
}

//การเล่นกับ pointer ด้วย struct
type student struct {
	name string
}

func setName(std student) {
	std.name = "Boy"
}
func setName2(std *student) {
	std.name = "Boy"
}
