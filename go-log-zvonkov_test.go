package main

import (
	"fmt"
	"testing"
)

func Testparse_time(t *testing.T) {
	s := "0:12"
	h, m := parse_time(s)
	fmt.Println(h)
	fmt.Println(m)
	if (h != "0") && (m != "12") {
		fmt.Println("Error in parse_time: ")
		fmt.Println(h)
		fmt.Println(m)
		t.Fatal("")
	}
}
