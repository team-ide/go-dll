package main

import "C"
import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	libPath := pwd + "/lib"

	lib, err := syscall.LoadDLL(libPath)
	if err != nil {
		panic(err)
	}
	validate, err := lib.FindProc("GetData")
	if err != nil {
		panic(err)
	}
	res1, res2, err := validate.Call()
	if err != nil {
		panic(err)
	}

	fmt.Println(res1)
	fmt.Println(res2)

}
