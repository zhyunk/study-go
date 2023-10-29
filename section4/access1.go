package section4

import (
	"fmt"
	"section4/lib2"
)

func Access1() {
	// 변수, 상수, 함수, 메서드 , 구조체 등 식별자
	// 대문자 : 패키지 외부에서 접근 가능
	// 소문자 : 패키지 외부 접근 불가 (패키지 내에서만 접근 가능)

	num1, num2 := 101, 999

	fmt.Println(num1, "은 100보다 큰 수? :", lib2.CheckNum1(num1))
	fmt.Println(num2, "은 1000보다 큰 수? :", lib2.CheckNum2(num2))
}

func init() {
	fmt.Println("Section4 Package - access1.go - go go")
}
