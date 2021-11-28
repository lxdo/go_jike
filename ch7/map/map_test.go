package _map

import "testing"

/**
map类型
*/
func TestInitMap(t *testing.T) {
	// 直接定义赋值
	//    map[键类型]值类型
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))

	// 初始化定义map
	m2 := map[int]int{}
	// 赋值
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))

	// make定义 map[键类型][值类型]  cap值
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

/**
判断map值是否存在
map元素的访问：在访问的key不存在时，仍会返回零值，不能通过返回nil来判断元素是否存在
*/
func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	// map值不存在 默认为0
	t.Log(m1[1]) // 0
	m1[2] = 0
	// map值被赋值为0
	t.Log(m1[2]) // 0
	// 判断map值是否存在
	// if  v(map值) ok(bool true值存在 false值不存在)
	if v, ok := m1[3]; ok {
		t.Logf("Key 3's value is %d", v)
	} else {
		t.Log("Key 3 is not existing.")
	}
}

/**
map 遍历
*/

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	// index value
	for i, v := range m1 {
		t.Log(i,v)
	}
}


