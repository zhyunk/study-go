package section5

import (
	"fmt"
	"unicode/utf8"
)

func String1() {
	str1 := "안녕하세용"

	// 바이트 수 구하기 : len(변수)
	fmt.Println(len(str1))

	// 실제 길이 구하기 1 : utf8.RuneCountInString(변수)
	fmt.Println(utf8.RuneCountInString(str1))

	// 실제 길이 구하기 2 : len([]rune( 변수 )) // 문자열을 배열에 저장해서 배열의 길이를 반환하는 형태
	fmt.Println(len([]rune(str1)))
}

func String2() {
	str1 := "Golang"
	str2 := "World"
	str3 := "고프로그래밍"

	// 다음과 같이 인덱스번호로 호출하면 ascii code 값이 나옴
	fmt.Println(str1, str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
	fmt.Println(str2, str2[0], str2[1], str2[2], str2[3], str2[4])
	fmt.Println(str3, str3[0], str3[1], str3[2], str3[3], str3[4], str3[5])

	// 다음과 같이 %c 형태로 출력하게 하면 ascii 문자가 나온다
	fmt.Printf("\n%s, %c %c %c %c %c %c", str1, str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
	fmt.Printf("\n%s, %c %c %c %c %c", str2, str2[0], str2[1], str2[2], str2[3], str2[4])
	fmt.Printf("\n%s, %c %c %c %c %c %c", str3, str3[0], str3[1], str3[2], str3[3], str3[4], str3[5])

	// 문자열을 rune 타입 배열로 만들면 인덱스로 각 문자에 접근 가능
	str4 := []rune(str3)
	fmt.Printf("\n%c, %c %c %c %c %c %c", str4, str4[0], str4[1], str4[2], str4[3], str4[4], str4[5])
}

func String3() {
	for i, c := range "Golang" {
		fmt.Printf("%c(%d)\t", c, i)
	}

	fmt.Printf("\n\n")

	for i, c := range "하이" {
		fmt.Printf("%c(%d)\t", c, i)
	}

	fmt.Printf("\n\n")

	for i, c := range []rune("하이") {
		fmt.Printf("%c(%d)\t", c, i)
	}
}
