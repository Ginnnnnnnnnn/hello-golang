package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Id    int      `json:"bianhao"`
	Name  string   `json:"mingzi"`
	Items []string `json:"mulu"`
}

func main() {

	// 结构体 -> json
	book1 := Book{
		Id:    1,
		Name:  "book1",
		Items: []string{"1", "2", "3"},
	}
	json1, _ := json.Marshal(book1)
	fmt.Println(string(json1))

	// json -> 结构体
	book2 := Book{}
	json.Unmarshal(json1, &book2)
	fmt.Println(book2)

}
