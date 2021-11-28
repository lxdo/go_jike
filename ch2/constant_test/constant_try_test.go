package consta_test

import "testing"

// 连续常量的简洁定义
// 从上向下递增
const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}

// 二进制连续位的常量简洁定义
// (没搞懂)
const (
	Readable = 1 << iota
	Writeable
	Executeable
)

func TestConstantTry1(t *testing.T) {
	a := 7 // 二进制 0111
	//a := 1 // 二进制 0001
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executeable == Executeable)
}

