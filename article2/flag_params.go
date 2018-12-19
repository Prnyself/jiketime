package main

import (
	"flag"
	"fmt"
)

//func init()  {
//
//}

func main() {
	name := flag.String("name", "everybody", "name to say hello")
	// 加了parse才可以在命令行中使用--help指令，显示参数列表
	// 并且flag.Parse()方法应该用在所有的flag参数存储载体的声明和设置之后
	flag.Parse()
	fmt.Println("hello", *name)
}
