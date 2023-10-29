package section3

import "fmt"

func Switch1() {
	// switch 다음에 오는 표현식 Expression 생략 가능
	// case 다음에 오는 표현식 expression 사용 가능
	// switch와 case의 들여쓰기 단계가 같다
	// 자동 break 때문에 fallthrouth 존재
	// 값이 아닌 변수 타입으로 분기 가능

	// 예제1
	a := -7
	switch {
	case a < 0:
		fmt.Println(a, "은(는) 음수")
	case a == 0:
		fmt.Println(a, "은(는) 0")
	case a > 0:
		fmt.Println(a, "은(는) 양수")
	}

	// 예제2
	switch b := 27; { // go에서 선호하는 방식
	case b < 0:
		fmt.Println(b, "은(는) 음수")
	case b == 0:
		fmt.Println(b, "은(는) 0")
	case b > 0:
		fmt.Println(b, "은(는) 양수")
	}

	// 예제3
	switch c := "go"; c { // 이런 식으로 작성하면 아래 case문에서 c 언급을 피할 수 있다.
	case "go":
		fmt.Println("Go !")
	case "java":
		fmt.Println("Java !")
	default:
		fmt.Println(c, " !")
	}

	// 예제4
	switch c := "go"; c + "lang" {
	case "go":
		fmt.Println("Go !")
	case "golang":
		fmt.Println("Golang !")
	case "java":
		fmt.Println("Java !")
	default:
		fmt.Println(c, " !")
	}

	// 예제5
	switch i, j := 20, 30; {
	case i < j:
		fmt.Println("i는 j보다 작다")
	case i == j:
		fmt.Println("i는 j와 같다")
	case i > j:
		fmt.Println("i는 j보다 크다")
	}

}

func init() {
	fmt.Println("Section3 Package - switch1.go - go go")
}
