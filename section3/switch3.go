package main

import "fmt"

func switch3() {
	// 예제 1
	a := 30 / 15
	switch a {
	case 2, 4, 6:
		fmt.Println("a는 7보다 작은 짝수인 ", a)
		if a == 2 {
			fmt.Println(a, "는 소수 중 유일한 짝수이다.")
		}
	case 1, 3, 5:
		fmt.Println("a는 7보다 작은 홀수인 ", a)
	}

	// 예제 2
	switch e := "go"; e {
	case "java":
		fmt.Println("자바!")
		fallthrough // 아래 조건이 맞지 않아도 해석함. 그렇기 때문에 마지막 case 문에는 사용할 수 없다.
	case "go":
		fmt.Println("고!")
		fallthrough
	case "js":
		fmt.Println("자바스크립트!")
		fallthrough
	case "php":
		fmt.Println("php !")
		fallthrough
	case "python":
		fmt.Println("파이썬 !")
		fallthrough
	case "kotlin":
		fmt.Println("코틀린 !")
	}
}
