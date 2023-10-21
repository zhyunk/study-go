package main

import "fmt"

func enumiota3() {

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
