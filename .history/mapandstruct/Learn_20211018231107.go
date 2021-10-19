package mapandstruct

import "fmt"

func Learn() {
	//Maps (key, value)
	//key = string , value = int
	//var m map[string]int
	m := map[string]int{"token": 10, "access": 20}

	fmt.Println(m)
	fmt.Println(m["access"])

	m["access"] = 200 //assign ค่าให้ map
	fmt.Println(m["access"])

	for _, v := range v {
		
	}
}
