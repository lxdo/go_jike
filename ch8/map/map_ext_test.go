package _map

import "testing"

/**
map 的value 可以是一个方法
与Go的Dock type接口方式一起，可以方便的实现单一方法对象的工厂模式
*/
func TestMapWithFunValue(t *testing.T) {
	// 定义 map 的键为int  value 为 func
	m := map[int]func(op int) int{}
	// 赋值
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

/**
实现set (没搞懂)
*/

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	// n := 1  1 is existing
	n := 3 // 3 is not existing
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}

	mySet[3] = true
	t.Log(len(mySet))
	// 删除
	delete(mySet, 1)
}
