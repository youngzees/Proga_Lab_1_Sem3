package structures

import (
	"fmt"
	"strconv"
	"strings"
)

type MyQueue struct {
	data []int
}

func (q *MyQueue) Push(val int) {
	q.data = append(q.data, val)
}

func (q *MyQueue) Pop() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("empty")
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v, nil
}

func (q *MyQueue) Serialize() string {
	strs := make([]string, len(q.data))
	for i, v := range q.data {
		strs[i] = strconv.Itoa(v)
	}
	return strings.Join(strs, " ")
}

func (q *MyQueue) Deserialize(str string) {
	q.data = nil
	parts := strings.Fields(str)
	for _, p := range parts {
		if v, err := strconv.Atoi(p); err == nil {
			q.data = append(q.data, v)
		}
	}
}

func (q *MyQueue) ToString() string {
	var sb strings.Builder
	sb.WriteString("[ front -> ")
	for _, v := range q.data {
		sb.WriteString(strconv.Itoa(v) + " ")
	}
	sb.WriteString("]")
	return sb.String()
}
