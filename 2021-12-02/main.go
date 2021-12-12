package main

import (
	"fmt"
	storage2 "github.com/spolischook/2021-12-02/storage"
	"github.com/spolischook/adventofcode/storage"
)

func main() {
	var provider storage.DataProvider
	provider = storage.FileProvider{}
	//provider = storage.HttpProvider{}
	data := storage2.SubmarineMoveCollection{}
	err := provider.GetData("/Users/spolischook/go/study/adventofcode/2021-12-02/input.txt", &data)
	if err != nil {
		panic(err)
	}

	sum := data.Sum()
	fmt.Println(sum)
	fmt.Println(sum.Multiply())
}
