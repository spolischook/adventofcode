package storage

import (
	"bufio"
	"os"
)

type FileProvider struct{}

func (p FileProvider) GetData(file string, data Data) (err error) {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f) // scanner's capacity for one line is 64K

	for scanner.Scan() {
		err = data.Append(scanner.Text())
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return err
}
