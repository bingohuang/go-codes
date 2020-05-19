package main

var b = 0

func test() func(int) int {
	a := 0
	return func(i int) int {
		a++
		b++
		return i + a + b
	}
}

func main() {
	f1 := test()
	f2 := test()

	println(f1(1))
	println(f2(1))
	println(f1(1))
	println(f2(1))
}
