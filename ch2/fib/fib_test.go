package fib

import (
	"testing"
)

// 变量赋值
// 赋值可以进行自动类型推断
// 在一个赋值语句中可以对多个变量进行同时赋值

var c int // 变量声明
func TestFibList(t *testing.T) {
	c = 1 // 声明变量赋值

	// 变量定义
	// var a int = 1
	// var b int = 1

	// 可简写为
	// var (
	//	 a int = 1
	//	 b     = 1 // int为默认类型 可以省略
	// )

	// 还可以简写为
	a := 1
	b := 1

	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

// 变量交换赋值

func TestExchange(t *testing.T) {
	a := 1
	b := 2

	//tmp:=a
	//a=b
	//b=tmp

	// 可以简写为
	a, b = b, a
	t.Log(a, b)
}
