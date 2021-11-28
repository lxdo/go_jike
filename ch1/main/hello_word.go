/**
go  文件的后缀是 .go

*/

//表示该hello.go文件所在的包是main ,在go中，每个文件都必须归属一个包
package main // 包，表明代码所在的模块（包）

import "fmt" // 引入代码依赖包

func main() {
	fmt.Println("hello world")
}

/**
通过 go build xxx.go 编译
通过 go run  xxx.go 直接执行
*/


