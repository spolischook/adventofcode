package storage

import (
	"errors"
	"strconv"
)

var EOC = errors.New("end of collection")

type IntCollection struct {
	pointer int // it should be uint, but I'm too lazy to convert it each iteration to int
	data    []int
}

func NewIntCollection(data []int) IntCollection {
	return IntCollection{
		data: data,
	}
}

func (c *IntCollection) GetWindow(l int) (int, error) {
	if c.pointer+l > len(c.data) {
		return 0, EOC
	}

	var sum int
	for i := 0; i < l; i++ {
		sum += c.data[c.pointer+i]
	}

	c.pointer++

	return sum, nil
}

func (c *IntCollection) Rewind() {
	c.pointer = 0
}

func (c *IntCollection) Append(el string) error {
	d, err := strconv.Atoi(el)
	if err != nil {
		return err
	}
	c.data = append(c.data, d)

	return nil
}

func (c IntCollection) Get() int {
	return c.data[c.pointer]
}

func (c *IntCollection) Next() bool {
	isNext := len(c.data) > (c.pointer + 1)

	if isNext == true {
		c.pointer++
		return true
	}

	return false
}
