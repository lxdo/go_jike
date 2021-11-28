package operator_test

import "testing"

// 数组比较
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b) // false 长度相同元素值不同
	//t.Log(a == c) 长度不同的数组无法比较
	t.Log(a == d) // true 长度相同元素值及顺序也相同
}

//位运算符
/**
&^ 按位置零
1 &^ 0 --1    对于左右两个操作数 右边的操作数位上是0 左边操作数位原来是什么就是什么
1 &^ 1 --0    对于左右两个操作数 右边的操作数位上是1 无论左边操作数位上是0还是1 结果都是0
0 &^ 1 --0
0 &^ 0 --0
*/
const (
	Readable = 1 << iota
	Writeable
	Executeable
)

func TestBitClear(t *testing.T) {
	a := 7 // 二进制 0111
	a = a &^ Readable
	a = a &^ Executeable
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executeable == Executeable)
}

