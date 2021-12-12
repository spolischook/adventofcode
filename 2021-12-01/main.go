package main

import (
	"fmt"
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

	// Pat 1
	count := 0
	prev, _ := data.Next()

	for current, err := data.Next(); err == nil; current, err = data.Next() {
		if prev < current {
			count++
		}

		prev = current
	}

	fmt.Printf("increase count: %d\n", count)

	// Part 2
	data.Rewind()
	count = 0
	const windowLen = 3
	prev, _ = data.GetWindow(windowLen)

	for current, err := data.GetWindow(windowLen); err == nil; current, err = data.GetWindow(windowLen) {
		if prev < current {
			count++
		}

		prev = current
	}

	fmt.Printf("increase window count: %d\n", count)
}
