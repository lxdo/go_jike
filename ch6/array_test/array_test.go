package array_test

import "testing"

/**
数组定义
*/
func TestArrayInit(t *testing.T) {
	// 定义数组 int 数组元素初始化默认0
	var arr [3]int
	t.Log(arr[0], arr[1], arr[2])

	// 定义数组 直接赋值

	arr1 := [4]int{1, 2, 3, 4}   // 指定数组元素个数
	arr3 := [...]int{1, 3, 4, 5} // 不指定数组元素个数

	arr1[1] = 5 // 数组元素重新赋值
	t.Log(arr1, arr3)
}

// 数组遍历
func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 3, 4, 5}
	// for循环
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}

	// go遍历
	for idx, e := range arr {
		t.Log(idx, e) // idx 索引  e 元素值
	}

	// 如果只需要数组元素值 用_占位
	for _, e := range arr {
		t.Log(e)
	}
}

/**
数组截取
a[开始索引(包含),结束索引(不包含)]
*/

func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	arr_sec1 := arr[:3]         // [1 2 3]
	arr_sec2 := arr[3:]         // [4 5]
	arr_sec3 := arr[1:2]        // [2]
	arr_sec4 := arr[1:3]        // [2,3]
	arr_sec5 := arr[1:len(arr)] // [2 3 4 5]
	t.Log(arr_sec1, arr_sec2, arr_sec3, arr_sec4, arr_sec5)
}

