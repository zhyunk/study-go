package section3

import "fmt"

func If2() {

	var a = 50
	b := 70

	// ex 1
	if a >= 65 {
		fmt.Println("65 이상")
	} else { // else 역시 컬리 브레이스가 else 옆에 있어야 정상 동작함
		fmt.Println("65 이하")
	}

	// ex 2
	if b >= 70 {
		fmt.Println("70 이상")
	} else {
		fmt.Println("70 이하")
	}

}
