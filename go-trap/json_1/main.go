package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := `{"age": 23, "name": "tony"}`
	m := map[string]interface{}{}
	_ = json.Unmarshal([]byte(s), &m)
	age := m["age"].(int) // panic: interface conversion: interface {} is float64, not int
	fmt.Println("age =", age)
}
