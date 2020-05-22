package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		switch i {
		case 0:
			fmt.Println(i)
		case 1, 2:
			fmt.Println(i)
			fallthrough
		case 3:
			fmt.Println(i)
		}
	}
}
