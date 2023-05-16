package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type configuration struct {
	Enabled bool
	Path    string
}

func main() {
	file, _ := os.Open("sample.json")
	defer file.Close()
	fmt.Println(file)
	decodeJson := json.NewDecoder(file)
	fmt.Println(decodeJson)
	conf := configuration{}
	err := decodeJson.Decode(&conf)

	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(conf.Path)
}
