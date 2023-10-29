package section2

import "fmt"

func Enumiota3() {

	const (
		_ = iota
		A
		_ // 생략 기능
		C
	)

	const (
		_ = iota + 0.75*2
		DEFAULT
		_
		GOLD
		PLATINUM
	)

	fmt.Println("A ", A)
	fmt.Println("C ", C)

	fmt.Println("DEFAULT ", DEFAULT)
	fmt.Println("GOLD ", GOLD)
	fmt.Println("PLATINUM ", PLATINUM)
}

func init() {
	fmt.Println("Section2 Package - enumiota3.go - go go")
}
