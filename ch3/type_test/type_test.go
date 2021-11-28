package type_test

import "testing"

// 基本数据类型
/*
	bool
	string
	int int8 int16 int32 int64
	uint uint8 uint16 uint32 uint64 uintptr  无符号整型
	bty  无符合8位整型的别名
	rune
	float32 float64
	complex64 complex128

*/

//	类型转化
/*
	1.go语言不允许隐式类型转换
	2.别名和原有类型也不能进行隐式类型转换
	3.类型转化必须显式转换
*/
type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

// 指针
// go不支持指针运算
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)

	t.Logf("%T %T", a, aPtr)
}

// string
func TestString(t *testing.T) {
	var s string // 值类型 默认初始化为""
	t.Log("*" + s + "*")
	t.Log(len(s))
	// 判断字符串是否为空
	if s==""{

	}
}

