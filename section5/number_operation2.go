package section5

import "fmt"

func No2() {
	var n1 uint8 = 125
	var n2 uint8 = 90

	fmt.Printf("\n%d + %d = %d", n1, n2, n1+n2)
	fmt.Printf("\n%d - %d = %d", n1, n2, n1-n2)
	fmt.Printf("\n%d * %d = %d", n1, n2, n1*n2)
	fmt.Printf("\n%d / %d = %d", n1, n2, n1/n2)
	fmt.Printf("\n%d %% %d = %d", n1, n2, n1%n2)
	fmt.Printf("\n%d << 2 = %d", n1, n1<<2)
	fmt.Printf("\n%d >> 2 = %d", n1, n1>>2)
	fmt.Printf("\n^%d = %d", n1, ^n1)

	var n3 int = 12
	var n4 float32 = 8.2
	var n5 uint16 = 1024
	var n6 uint32 = 120000

	fmt.Printf("\n%d + %f = %f", n3, n4, float32(n3)+n4)
	fmt.Printf("\n%d + %f = %d", n3, n4, n3+int(n4))
	fmt.Printf("\n%d + %d = %d", n5, n6, n5+uint16(n6))
}
