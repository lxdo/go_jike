package slice_test

import "testing"

/**
切片
	内部结构   len:元素的个数
              cap: 内部数组的容量
*/
func TestSliceInit(t *testing.T) {
	// 声明切片
	var s0 []int
	t.Log(len(s0), cap(s0))
	// 向切片中填充元素 append
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	// 声明直接赋值
	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))
	// make 函数声明
	// 参数 make(类型 ,len， cap)
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	// len为可访问的元素个数
	// t.Log(s2[0], s2[1], s2[2],s2[3],s2[4])  报错 访问元素超出len
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
}

/**
切片实现可变长
*/
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		// cap成倍增长
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

/**
切片共享内存
*/
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	// cap 从截取元素Apr一直到最后
	t.Log(Q2, len(Q2), cap(Q2)) // [Apr May Jun] 3 9
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer)) // [Jun Jul Aug] 3 7
	// 修改summer Q2也受影响 他们共享同一个内存
	summer[0] = "Unknow"
	t.Log(Q2)   // [Apr May Unknow]
	t.Log(year) // [Jan Feb Mar Apr May Unknow Jul Aug Sep Oct Nov Dec]
}

/**
数组和切片异同

数组容量不可伸缩   切片可以
数组相同维数，相同长度可以比较   切片不可比较

*/

