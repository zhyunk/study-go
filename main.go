package main

import (
	"fmt"
	_ "section1"
	_ "section2"
	_ "section3"
	s4 "section4"
)

func main() {
	fmt.Println("Root Package - Main Method go go")

	// s1.Test()
	// s2.Enumiota3()
	// s3.For3()
	s4.Init2()
}

func init() {
	fmt.Println("Root Package - Init Method go go")
}
