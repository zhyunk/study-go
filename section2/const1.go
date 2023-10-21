package main

import "fmt"

func const1() {
	// 상수
	// const 사용 초기화, 한 번 선언 후 값 변경 금지
	// 고정 된 값 관리용도

	const a string = "Test1"
	const b = "Test2"
	const c int32 = 10 * 10
	const e = 35.6
	const f = false
	const t = true

	/*
		에러 발생하는 경우

		const g string
		g = "Test3"

		const d = getHeight() // 함수는 항상 같은 값을 반환한다는 보장이 없기 때문에 불가능한 형태이다. 실행시 에러남

	*/

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)
	fmt.Println("e : ", e)
	fmt.Println("f : ", f)
	fmt.Println("t : ", t)
}
