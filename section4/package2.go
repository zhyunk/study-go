package section4

import (
	"fmt"
	"section4/lib"
)

func Package2() {
	fmt.Println("10보다 큰 수? : ", lib.CheckNum(109))
}

func init() {
	fmt.Println("Section4 Package - package2.go - go go")
}
