package main

import "C"
import (
	"fmt"
	"unsafe"
)

//export GetData
func GetData() *int {
	data := "这是一段中文文本，Golang 字符串编码是 utf-8，在 C 语言里需要转换成 GBK 编码才能正确显示。"
	println("GetData:", data)
	fmt.Println("GetData:", data)
	res := (*int)(unsafe.Pointer(C.CString(data)))
	return res
}

//export SetData
func SetData(data string) {
	fmt.Println("SetData:", data)
}

func main() {
}
