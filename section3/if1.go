package section3

import "fmt"

func If1() {
	// 조건문
	// 반드시 bool 검사 -> 1, 0 사용 불가 (자동 형 변환 안됨)
	// 소괄소 사용 x

	var a = 20
	b := 20

	// 예제 1
	if a >= 15 {
		fmt.Println("15 이상")
	}

	// 예제 2
	if b >= 25 { // 컬리 브레이스는 조건식 옆에 있어야 실행된다.
		fmt.Println("25 이상")
	}

	/*
		// 에러 발생 1
		if b >= 25
		{
			fmt.Print("err")
		}


		// 에러 발생 2 : 1줄 이하의 코드를 가졌을 때도 컬리 브레이스를 꼭 명시해주어야 함
		if b >= 5
			fmt.Print("25 이상")


		// 에러 발생 3 : 정확하게 true, false 가 아닌 결과를 내는 조건식이 들어오면 에러가 남
		if c := 1; c {
			fmt.Print("True")
		}
	*/

	// 예제 3
	if c := 40; c >= 35 {
		fmt.Println("35 이상")
	}

}
