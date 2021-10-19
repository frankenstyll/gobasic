package mapandstruct

import (
	"fmt"
)

func Learn() {
	//Maps (key, value)
	//key = string , value = int
	//var m map[string]int
	m := map[string]int{"token": 10, "access": 20}

	fmt.Println(m)
	fmt.Println(m["access"])

	m["access"] = 200 //assign ค่าให้ map
	fmt.Println(m["access"])

	for key, v := range m {
		fmt.Printf("%v => %v \n", key, v)
	}

	//delete map ด้วย key
	delete(m, "access")
	fmt.Println(m)
	//add map
	m["username"] = 123
	fmt.Println(m)

	m2 := make(map[string[]]string)
	m2["t"] = "index1"
	fmt.Println(m2)

	//*****************************struct ใช้บ่อยมากก
}