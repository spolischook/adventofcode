package storage

import (
	"bufio"
	"os"
	"strconv"
)

type FileProvider struct{}

func (p FileProvider) GetData() (data IntCollection, err error) {
	f, err := os.Open("/Users/spolischook/go/study/adventofcode/day1/data.txt") // get file path from args
	if err != nil {
		return IntCollection{}, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f) // scanner's capacity for one line is 64K

	for scanner.Scan() {
		d, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}
		data.Append(d)
	}

	if err := scanner.Err(); err != nil {
		return IntCollection{}, err
	}

	return data, err
}
