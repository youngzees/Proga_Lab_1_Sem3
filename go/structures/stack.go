package structures

import (
	"fmt"
	"strconv"
	"strings"
)

type MyStack struct {
	data []int
}

func (s *MyStack) Push(val int) {
	s.data = append(s.data, val)
}

func (s *MyStack) Pop() (int, error) {
	if len(s.data) == 0 {
		return 0, fmt.Errorf("empty")
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, nil
}

func (s *MyStack) Serialize() string {
	strs := make([]string, len(s.data))
	for i, v := range s.data {
		strs[i] = strconv.Itoa(v)
	}
	return strings.Join(strs, " ")
}

func (s *MyStack) Deserialize(str string) {
	s.data = nil
	parts := strings.Fields(str)
	for _, p := range parts {
		if v, err := strconv.Atoi(p); err == nil {
			s.data = append(s.data, v)
		}
	}
}

func (s *MyStack) ToString() string {
	var sb strings.Builder
	sb.WriteString("[ top: ")
	for i := len(s.data) - 1; i >= 0; i-- {
		sb.WriteString(strconv.Itoa(s.data[i]) + " ")
	}
	sb.WriteString("]")
	return sb.String()
}
