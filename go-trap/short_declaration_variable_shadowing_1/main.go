package main

func main() {
	var a int = 13
	println(a, &a)          // 13 0xc000044770
	a, d := 23, "hello, go" // 注：此处只是对a重新赋值
	println(&a)             // 0xc000044770 地址相同
	println(a, d)           // 23 hello, go

	if a == 23 {
		var a int = 33 // 注：遮蔽了main函数中声明的变量a
		println(a, &a) // 33 0xc000044768 地址不同
	}
}
