package main

import (
	"flag"
	"fmt"
	"strings"
)

//注册命令行参数
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	//解析命令行参数 并对已经注册的命令行参数进行匹配，如果匹配到了就赋值
	//没匹配到的就放在flag.Args()中
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
