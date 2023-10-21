package main

import "fmt"

func var3() {
	// 짧은 선언
	// 반드시 함수 안에서만 사용해야 된다.
	// 선언 후 할당하면 에러가 발생한다.
	// > 마치 자바의 상수 느낌이군
	// 주로 제한된 범위의 함수 내에서 사용 할 경우 코드 가독성을 높일 수 있다.

	shortVar1 := 3
	shortVar2 := "Test"
	shortVar3 := true

	// shortVar1 := 10 // 예외 발생

	fmt.Println("shortVar1 : ", shortVar1)
	fmt.Println("shortVar2 : ", shortVar2)
	fmt.Println("shortVar3 : ", shortVar3)

	// 사용 예시
	if i := 10; i < 11 {
		fmt.Println("Short Variable Test Success!")
		i++
		fmt.Println("i = ", i) // i = 11
	}
}
