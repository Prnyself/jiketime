package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type intList []int

func (iList *intList) String() string {
	return fmt.Sprint(*iList)
}

// 每一个flag都会调用Set方法
// 比如-i=1 -i=2, 就会调用两次Set方法, 因为是指针, 所以*intList的长度会增加
// 默认情况下, 后写的参数会代替前一个如: -name=a -name=b, 实际取到的name值是b
func (iList *intList) Set(value string) error {
	// len(*iList) 用于判断该参数是否有多个，如果已经存在则会返回errors
	if len(*iList) > 0 {
		return errors.New("intList flag already set")
	}
	// 这个for循环是为了适配使用逗号分隔传参数的情况, 如-i=1,2,3
	// 利用strings包的Split参数进行分割, 并添加到*iList后
	for _, dt := range strings.Split(value, ",") {
		num, err := strconv.Atoi(dt)
		if err != nil {
			return err
		}
		*iList = append(*iList, num)
	}
	return nil
}

// 常用init方法来读取传入的参数
//func init()  {
//
//}

func main() {
	var iList intList
	var yesOrNo bool
	// 也可以用flag.StringVar()方法，该方法接收第一个参数的地址来存值，无返回值
	name := flag.String("name", "everybody", "name to say hello")
	dr := flag.Duration("time", 2*time.Second, "time duration")
	// 如果是bool类型的, 可以不用输入值, 比如: 有-test, 就会被解析为true
	flag.BoolVar(&yesOrNo, "test", false, "test mode or not")
	flag.Var(&iList, "i", "Put i into iList")

	// 重写flag.Usage方法可以修改--help时的显示
	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stdout, "Usage of %v:\n", os.Args[0])
		flag.PrintDefaults()
	}

	// 加了parse才可以在命令行中使用--help指令，显示参数列表
	// 并且flag.Parse()方法应该用在所有的flag参数存储载体的声明和设置之后
	// 常放在main函数的最初, 不能放到init方法里
	flag.Parse()
	fmt.Println("hello", *name)
	fmt.Println(iList)
	fmt.Println("time duration:", *dr)
	fmt.Println("test mode?", yesOrNo)
	// output:
	//hello everybody
	//[]
	//time duration: 2s
	//test mode? false
}
