package loop

import "testing"

// go用for循环实现while循环
func TestWhileLoop(t *testing.T) {
	n := 0
	/*
	  等同于 while (n<5)
	 */
	for n < 5 {
		t.Log(n)
		n++
	}
}

