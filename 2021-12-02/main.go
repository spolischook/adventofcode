package main

import (
	"github.com/spolischook/adventofcode/storage"
)

func main() {
	var provider storage.DataProvider
	provider = storage.FileProvider{}
	//provider = storage.HttpProvider{}
	data, err := provider.GetData()
	if err != nil {
		panic(err)
	}
}
