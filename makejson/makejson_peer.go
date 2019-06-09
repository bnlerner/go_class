package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := make(map[string]string)
	var name string
	fmt.Scan(&name)
	var address string
	fmt.Scan(&address)
	m["name"] = name
	m["address"] = address
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json marshal fail:", err)
	}
	fmt.Println(string(b))
}
