package section3

import "fmt"

func If3() {
	i := 44

	// if - else if 예제 1
	if i >= 120 {
		fmt.Println("120 이상")
	} else if i >= 100 && i < 120 {
		fmt.Println("100 이상 120 미만")
	} else if i < 100 && i >= 50 {
		fmt.Println("50 이상 100 미만")
	} else {
		fmt.Println("50 미만")
	}
}

func init() {
	fmt.Println("Section3 Package - if3.go - go go")
}
