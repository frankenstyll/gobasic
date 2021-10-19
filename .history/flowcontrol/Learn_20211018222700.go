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

	b, err = ioutil.ReadFile("file.txt")
	if err 
}
