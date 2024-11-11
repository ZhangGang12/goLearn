package main

import (
	"fmt"
	"os"
	"strings"
)

func echoV1() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echoV2() {
	var s string
	for k, v := range os.Args {
		s += fmt.Sprintf(" %d %s", k, v)
	}
	fmt.Println(s)
}

func echoV3() {
	a := strings.Join(os.Args, " ")
	_ = a
}

func echoV4() {
	a := os.Args[1:]
	_ = a
}
