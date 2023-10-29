package section4

import (
	"fmt"
	"section4/lib"
)

var num int32

func init() {
	num = 30
	fmt.Println("Section4 Package - init2.go - go go")
}

func Init2() {
	fmt.Println(num, "> 10 :", lib.CheckNum(int(num)))
}
