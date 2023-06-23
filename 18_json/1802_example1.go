/* Json with structs,maps */

package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// structs to json
	type Book struct {
		Title  string
		Author string
		Year   int
	}

	myBook := Book{"Hello Golang", "John Mike", 2021}
	bytes_1, _ := json.Marshal(myBook)
	fmt.Println(string(bytes_1))

	// json to structs
	jsonInput1 := `{"Title": "Hello Golang", "Author":"John Mike", "Year": 2021}`

	res_struct := Book{}
	err1 := json.Unmarshal([]byte(jsonInput1), &res_struct)
	if err1 != nil {
		fmt.Println("JSON decode error!")
		return
	}
	fmt.Println(res_struct)

	// maps to json
	fileCount := map[string]int{
		"cpp": 10,
		"js":  8,
		"go":  10,
	}
	bytes_2, _ := json.Marshal(fileCount)
	fmt.Println(string(bytes_2))

	// json to map
	jsonInput2 := `{
        "apples": 10,
        "mangos": 20,
        "grapes": 20
    }`

	var fruitBasket map[string]int
	err := json.Unmarshal([]byte(jsonInput2), &fruitBasket)

	if err != nil {
		fmt.Println("JSON decode error!")
		return
	}

	fmt.Println(fruitBasket)
}
