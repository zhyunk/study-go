package section3

import (
	"fmt"
	"math/rand"
	"time"
)

func Switch2() {
	// 랜덤시드를 유닉스 시간 단위로 설정해 주면 매번 다른 난수를 얻을 수 있다.
	// 더 안전한(예측하기 힘든) 난수를 얻고 싶다면 crypt/rand 패키지를 사용해야 한다. > https://www.joinc.co.kr/w/GoLang/example/randomnumber
	rand.New(rand.NewSource(time.Now().UnixNano())) // rand.Seed(time.Now().UnixNano()) 가 deprecated 됨

	switch i := rand.Intn(100); { // rand.Intn(100); // 0 ~ 99 범위에서 임의의 수 생성
	case i >= 50 && i < 100:
		fmt.Println("i -> ", i, " , 50 이상 100 미만")
	case i >= 25 && i < 50:
		fmt.Println("i -> ", i, " , 25 이상 50 미만")
	default:
		fmt.Println("i -> ", i, " , 25 미만")
	}
}

func init() {
	fmt.Println("Section3 Package - switch2.go - go go")
}
