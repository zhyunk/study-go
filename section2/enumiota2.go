package section2

import "fmt"

func Enumiota2() {

	const (
		A = iota
		B
		C
	)

	const (
		D = iota * 10
		E
		F
	)

	const (
		Jan = iota + 1
		Feb
		Mar
		Apr
		May
		Jun
	)

	fmt.Println("A ", A)
	fmt.Println("B ", B)
	fmt.Println("C ", C)

	fmt.Println("D ", D)
	fmt.Println("E ", E)
	fmt.Println("F ", F)

	fmt.Println("Jan ", Jan)
	fmt.Println("Feb ", Feb)
	fmt.Println("Mar ", Mar)
	fmt.Println("Apr ", Apr)
	fmt.Println("May ", May)
	fmt.Println("Jun ", Jun)
}

func init() {
	fmt.Println("Section2 Package - enumiota2.go - go go")
}
