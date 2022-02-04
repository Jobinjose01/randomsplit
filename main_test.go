package main

import (
	"fmt"
	"testing"
)

func TestSplitMyNumber(t *testing.T) {

	results := SplitMyNumber(167, 3)
	fmt.Println(results)

	results = SplitMyNumber(11, 10)
	fmt.Println(results)

	results = SplitMyNumber(11, 5)
	fmt.Println(results)

	results = SplitMyNumber(1600, 50)
	fmt.Println(results)

	//Execute Same Test Cases 100 times
	fmt.Println("Execute same test case 100 times!")
	for i := 0; i < 100; i++ {
		results := SplitMyNumber(167, 3)
		fmt.Println(results)
	}
}

func BenchmarkSplitMyNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		results := SplitMyNumber(167, 3)
		fmt.Println(results)
	}
}
