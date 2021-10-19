package flowcontrol

import (
	"fmt"
	"io/ioutil"
	"runtime"
)

func Learn() {

	//if
	score := 10
	if score == 10 {
		fmt.Println("score = 10")
	} else if score == 20 {
		fmt.Println("score = 20")
	} else {
		fmt.Println("not match")
	}

	//switch
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("MacOS")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Printf("Other %s \n", os)
	}

	//ioutil.readfile คืออ่านไฟล์
	//_, err := ioutil.ReadFile("file.txt")  => _ คือละเว้นว่าจะไม่ใช้ เพราะ ioutil.readfile จะ return มา 2 ค่า

	_, err := ioutil.ReadFile("file.txt")
	if err != nil {
		//ถ้า error ตัวแปร err จะได้ค่า nil
		fmt.Println(err)
	}
}
