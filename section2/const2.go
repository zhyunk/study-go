package section2

import "fmt"

func Const2() {
	const a, b int = 10, 99
	const c, d, e = true, 0.84, "test"

	const (
		f    = 3.14
		g, z = "gim", "zhyun"
	)

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)
	fmt.Println("e : ", e)
	fmt.Println("d : ", d)
	fmt.Println("f : ", f)
	fmt.Println("g : ", g)
	fmt.Println("z : ", z)
}
