package section3

import "fmt"

func For1() {
	// go에서 유일하게 제공되는 반복문
	// 다양한 사용법 숙지

	// 예제 1
	for i := 1; i < 5; i++ {
		fmt.Println("ex1 : ", i)
	}

	/*
		// 에러 발생 1
		for i := 0; i < 5; i++
		{
			fmt.Println("ex1 : ", i)
		}

		// 에러 발생 2
		for i := 0; i < 5; i++
			fmt.Println("ex1 : ", i)

	*/

	// 예제 2 : 무한 루프
	roopI := 0
	for {
		if roopI == 5 {
			break
		}

		roopI++
		fmt.Println(roopI)
	}

	// 예제 3 : Range 용법
	loc := []string{"Seoul", "Busan", "Incheon"}

	// 인덱스와 값 출력
	for index, name := range loc {
		fmt.Println("ex3 : ", index, name)
	}
	// 인덱스만 출력
	for index := range loc {
		fmt.Println("ex3 : ", index)
	}
	// 값만 출력
	for _, name := range loc {
		fmt.Println("ex3 : ", name)
	}
}

func init() {
	fmt.Println("Section3 Package - for1.go - go go")
}
