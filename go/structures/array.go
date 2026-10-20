package structures

import (
	"fmt"
	"strconv"
	"strings"
)

type MyArray struct {
	data []int
}

func (a *MyArray) Push(val int) {
	a.data = append(a.data, val)
}

func (a *MyArray) Del(index int) error {
	if index < 0 || index >= len(a.data) {
		return fmt.Errorf("bad index")
	}
	a.data = append(a.data[:index], a.data[index+1:]...)
	return nil
}

func (a *MyArray) Get(index int) (int, error) {
	if index < 0 || index >= len(a.data) {
		return 0, fmt.Errorf("bad index")
	}
	return a.data[index], nil
}

func (a *MyArray) Set(index, val int) error {
	if index < 0 || index >= len(a.data) {
		return fmt.Errorf("bad index")
	}
	a.data[index] = val
	return nil
}

func (a *MyArray) Length() int {
	return len(a.data)
}

func (a *MyArray) Serialize() string {
	strs := make([]string, len(a.data))
	for i, v := range a.data {
		strs[i] = strconv.Itoa(v)
	}
	return strings.Join(strs, " ")
}

func (a *MyArray) Deserialize(s string) {
	a.data = nil
	parts := strings.Fields(s)
	for _, p := range parts {
		if v, err := strconv.Atoi(p); err == nil {
			a.data = append(a.data, v)
		}
	}
}
