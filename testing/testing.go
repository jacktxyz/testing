package testing

import (
	"encoding/json"
	"fmt"
)

// testing map
func TestingM() {
	m := map[string]int{}

	datas := []string{"string", "int", "float", "struct", "type", "func", "main", "testing", "fmt"}

	for i, data := range datas {
		m[data] = i
	}

	fmt.Println(m)

	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

// testing slice
func TestingS() {
	s := make([]string, 0)

	s = append(s, "1")
	s = append(s, "2")
	s = append(s, "3")
	s = append(s, "4")
	s = append(s, "5")

	fmt.Println(s)
}
