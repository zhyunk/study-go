package section2

import "fmt"

func Var2() {
	var (
		name      string = "machine"
		height    int32
		weight    float32
		isRunning bool
	)

	height = 250
	weight = 350.56
	isRunning = true

	fmt.Println("name : ", name)
	fmt.Println("height : ", height)
	fmt.Println("weight : ", weight)
	fmt.Println("isRunning : ", isRunning)
}

func init() {
	fmt.Println("Section2 Package - var2.go - go go")
}
