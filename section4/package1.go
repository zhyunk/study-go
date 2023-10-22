package section4

import (
	"fmt"
	"os"
)

func Package1() {
	// 패키지 이름 = 디렉토리 이름
	// 같은 패키지 내 -> 소스 파일들은 디렉토리 명을 패키지 명으로 사용한다.
	// 네이밍 규칙 : 소문자 = private , 대문자 = public
	// go :  컴파일러는 main 패키지를 프로그램의 시작점 start point로 인식
	// main 함수 외의 파일은 자신이 포함 된 폴더 이름을 패키지명으로 갖는다.

	var name string

	fmt.Println("이름은 ? : ")
	fmt.Scanf("%s", &name)

	fmt.Fprintf(os.Stdout, "Hi! %s\n", name)
}
