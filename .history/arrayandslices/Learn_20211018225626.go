package arrayandslices

import (
	"fmt"

	"honnef.co/go/tools/printf"
)

func Learn() {

	//ประกาศ array
	var arr [2]string
	arr[0] = "A"
	arr[1] = "B"

	fmt.Println(arr)

	//ประกาศค่าเริ่มต้น
	arr2 := [6]int{1, 2, 3, 9, 8, 7}
	for _, v := range arr2 {
		fmt.Println(v)
	}

	fmt.Println("-------------------------------")

	//slice คือ build on top ของ array
	//คล้ายๆ substring [index เริ่ม: index สุดท้ายที่เอา]
	var s []int = arr2[2:4]
	fmt.Println(s)

	//กำหนดค่าเริ่มต้นให้ slice
	s2 := []int{2, 3, 5, 7, 11, 13}
	s2[0] = 22 // set ค่าให้ slice index 0
	fmt.Println(s2[0])
	fmt.Printf("len=%d cap=%d \n", len(s2), cap)
}
