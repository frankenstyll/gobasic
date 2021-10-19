package arrayandslices

import "fmt"

func Learn() {

	//array
	var arr [2]string
	arr[0] = "A"
	arr[1] = "B"

	fmt.Println(arr)

	//ประกาศค่าเริ่มต้น
	arr2 := [3]int{1, 2, 3}
	for _, v := range arr2 {
		fmt.Println(v)
	}

}
