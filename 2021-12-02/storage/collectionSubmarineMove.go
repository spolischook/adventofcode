package storage

import (
	"fmt"
	"strconv"
	"strings"
)

type SubmarineMoveCollection struct {
	pointer int
	data    []SubmarineMove
}

type SubmarineMove struct {
	x int
	z int
}

func (c SubmarineMoveCollection) Aim() SubmarineMove {
	var aim int
	sm := SubmarineMove{}
	calc := func(sm *SubmarineMove, sm2 SubmarineMove) {
		aim += sm2.z
		sm.x += sm2.x
		sm.z += aim*sm2.x
	}
	calc(&sm, c.Get())
	for c.Next() {
		calc(&sm, c.Get())
	}

	return sm
}

func (c SubmarineMoveCollection) Sum() SubmarineMove {
	m := c.Get()
	for c.Next() {
		m.x += c.Get().x
		m.z += c.Get().z
	}

	return m
}

func (c SubmarineMoveCollection) Get() SubmarineMove {
	return c.data[c.pointer]
}

func (c *SubmarineMoveCollection) Next() bool {
	isNext := len(c.data) > (c.pointer + 1)

	if isNext == true {
		c.pointer++
		return true
	}

	return false
}
func (c *SubmarineMoveCollection) Append(el string) error {
	parts := strings.Split(el, " ")
	if len(parts) != 2 {
		return fmt.Errorf(`unexpected element "%s"`, el)
	}

	n, err := strconv.Atoi(parts[1])
	if err != nil {
		return err
	}
	sm := SubmarineMove{}

	err = sm.Move(parts[0], n)
	if err != nil {
		return err
	}

	c.data = append(c.data, sm)

	return nil
}

func (sm *SubmarineMove) Move(dir string, n int) error {
	switch dir {
	case "down":
		sm.Down(n)
	case "up":
		sm.Up(n)
	case "forward":
		sm.Forward(n)
	default:
		return fmt.Errorf(`unknown "%s" direction`, dir)
	}

	return nil
}

func (sm *SubmarineMove) Down(n int) {
	sm.z += n
}
func (sm *SubmarineMove) Up(n int) {
	sm.z -= n
}
func (sm *SubmarineMove) Forward(n int) {
	sm.x += n
}
func (sm SubmarineMove) Multiply() int {
	return sm.x*sm.z
}
