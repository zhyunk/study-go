package section2

import "fmt"

func Var1() {
	// 기본 초기화
	// 정수 타입 : 0
	// 실수 타입 : 0.0
	// 문자열 : ""
	// 논리 : true, false
	// 변수명 : 숫자 첫글자x , 대소문자 구분 , 문자 숫자 밑줄 특수문자 사용 가능
	// 변수 및 상수 : 함수 내외 사용 가능

	// 변수 : var 변수명 자료형 = 초기화
	var a int = 4
	var b string = "asd"
	var c, d, e int
	var f, g, h int = 1, 2, 3
	var i float32 = 11.4
	var j string = "go lang"
	var k = 4.74 // 선언 동시 초기화
	var l = "Hi 고!"
	var m = true

	c = 11
	d = 21
	e = 31

	fmt.Println("a : ", a,
		"\nb : ", b,
		"\nc : ", c,
		"\nd : ", d,
		"\ne : ", e,
		"\nf : ", f,
		"\ng : ", g,
		"\nh : ", h,
		"\ni : ", i,
		"\nj : ", j,
		"\nk : ", k,
		"\nl : ", l,
		"\nm : ", m)
}

func init() {
	fmt.Println("Section2 Package - var1.go - go go")
}
