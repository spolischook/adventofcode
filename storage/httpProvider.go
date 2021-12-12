package storage

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://adventofcode.com/2021/day/1/input"

type HttpProvider struct{}
func (p HttpProvider) GetData() (IntCollection, error) {
	resp, err := http.Get(url)
	if err != nil {
		return IntCollection{}, err
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return IntCollection{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return IntCollection{}, fmt.Errorf(`unexpected %d status code and message: "%s"`, resp.StatusCode, data)
	}

	dataStr := string(data)
	fmt.Println(dataStr)

	return IntCollection{}, nil
}
