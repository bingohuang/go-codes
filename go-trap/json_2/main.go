package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	s := `{"age": 23, "name": "tony"}`
	m := map[string]interface{}{}

	d := json.NewDecoder(strings.NewReader(s))
	d.UseNumber()

	_ = d.Decode(&m)
	age, _ := m["age"].(json.Number).Int64()
	fmt.Println("age =", age) // age = 23
}
