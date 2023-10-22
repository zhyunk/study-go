package section3

import "fmt"

func For3() {
	// 예제 1
Loop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break Loop
			}
			fmt.Println("ex1 :", i, j)
		}
	}

	// 예제 2
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("ex2 :", i)
	}

	// 예제 3
here:
	// 에러 발생 : 레이블 밑에 관련 없는 소스코드(println 등)가 오면 에러남
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue here
			}
			fmt.Println("ex3 :", i, j)
		}
	}

}
