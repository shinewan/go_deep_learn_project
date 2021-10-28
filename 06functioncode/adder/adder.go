package main

import (
	"fmt"
)

/**
	“正统” 函数式编程
	1.不可变性：不能有状态(举例：不能有变量等)，只有常量和函数
	2.函数只能有一个参数
 */


func adder() func(int) int {
	sum := 0
	//sum对于匿名函数来说，就是外部变量
	return func(v int) int{
		sum += v
		return sum
	}
}

//声明函数类型
type iAdder func(int) (int,iAdder)
//写一个方法，返回值是自定义的函数
func adder2(base int) iAdder{//这个函数的参数是base  返回值是 func(v int)

	//对于匿名函数func来说，base就是外部变量
	return func(v int) (int,iAdder){//这个函数的参数是v
		return base + v ,adder2(base+v) //返回的时候，又调用了一次adder2
	}
}

func main() {
	a := adder()
	for i:= 0;i < 10 ;i++ {
		fmt.Printf("0+...+%d = %d\n",i,a(i))
	}
	fmt.Println("----------------")
	b := adder2(0)
	for i:= 0;i < 10 ;i++ {
		var s int
		s,b = b(i)
		fmt.Printf("0+...+%d = %d\n",i,s)
	}
}
