package section4

import (
	"fmt"
	check10Up "section4/lib"
	_ "section4/lib2" // 바로 사용하지 않는 경우 _ 를 붙여두면 지금 사용하지 않았어도 삭제되지 않는다.
	check100Up "section4/lib2"
)

func Access2() {
	var num = 109

	fmt.Println(num, "> 10 :", check10Up.CheckNum(num))
	fmt.Println(num, "> 100 :", check100Up.CheckNum1(num))
	fmt.Println(num, "> 1000 :", check100Up.CheckNum2(num))
}
