package section5

import (
	"fmt"
	"math"
)

func No1() {
	var n1 uint8 = math.MaxUint8
	var n2 uint16 = math.MaxUint16
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64

	fmt.Println("MaxUint8 : ", n1)
	fmt.Println("MaxUint16 : ", n2)
	fmt.Println("MaxUint32 : ", n3)
	fmt.Println("MaxUint64 : ", n4)
	fmt.Println("MaxFloat32 : ", math.MaxFloat32)
	fmt.Println("MaxFloat64 : ", math.MaxFloat64)

	n5 := 100000       // int 자료형으로 생성됨
	n6 := int16(10000) // int16 자료형으로 형변환 후 생성됨
	n7 := uint8(100)   // int8 자료형으로 형변환 후 생성됨

	// fmt.Println(n5 + n6) // 다른 타입끼리의 연산은 예외를 발생시킨다.
	fmt.Println(n5 + int(n6))
	fmt.Println(n5 + int(n7))

	fmt.Println(n6 > int16(n7))
	fmt.Println(n6-int16(n7) > 5000)

}
