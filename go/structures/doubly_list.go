package structures

import (
	"fmt"
	"strconv"
	"strings"
)

type NodeL struct {
	Val  int
	Prev *NodeL
	Next *NodeL
}

type DoublyList struct {
	Head *NodeL
	Tail *NodeL
}

func (l *DoublyList) PushHead(val int) {
	n := &NodeL{Val: val, Next: l.Head}
	if l.Head != nil {
		l.Head.Prev = n
	} else {
		l.Tail = n
	}
	l.Head = n
}

func (l *DoublyList) PushTail(val int) {
	n := &NodeL{Val: val, Prev: l.Tail}
	if l.Tail != nil {
		l.Tail.Next = n
	} else {
		l.Head = n
	}
	l.Tail = n
}

func (l *DoublyList) PushBefore(target, val int) error {
	t := l.Head
	for t != nil && t.Val != target {
		t = t.Next
	}
	if t == nil {
		return fmt.Errorf("target not found")
	}
	n := &NodeL{Val: val, Next: t, Prev: t.Prev}
	if t.Prev != nil {
		t.Prev.Next = n
	} else {
		l.Head = n
	}
	t.Prev = n
	return nil
}

func (l *DoublyList) PushAfter(target, val int) error {
	t := l.Head
	for t != nil && t.Val != target {
		t = t.Next
	}
	if t == nil {
		return fmt.Errorf("target not found")
	}
	n := &NodeL{Val: val, Prev: t, Next: t.Next}
	if t.Next != nil {
		t.Next.Prev = n
	} else {
		l.Tail = n
	}
	t.Next = n
	return nil
}

func (l *DoublyList) PopHead() error {
	if l.Head == nil {
		return fmt.Errorf("empty")
	}
	l.Head = l.Head.Next
	if l.Head != nil {
		l.Head.Prev = nil
	} else {
		l.Tail = nil
	}
	return nil
}

func (l *DoublyList) PopTail() error {
	if l.Tail == nil {
		return fmt.Errorf("empty")
	}
	l.Tail = l.Tail.Prev
	if l.Tail != nil {
		l.Tail.Next = nil
	} else {
		l.Head = nil
	}
	return nil
}

func (l *DoublyList) Del(val int) error {
	for t := l.Head; t != nil; t = t.Next {
		if t.Val == val {
			if t.Prev != nil {
				t.Prev.Next = t.Next
			} else {
				l.Head = t.Next
			}
			if t.Next != nil {
				t.Next.Prev = t.Prev
			} else {
				l.Tail = t.Prev
			}
			return nil
		}
	}
	return fmt.Errorf("not found")
}

func (l *DoublyList) DelIndex(index int) error {
	if index < 0 {
		return fmt.Errorf("bad index")
	}
	t := l.Head
	for i := 0; i < index && t != nil; i++ {
		t = t.Next
	}
	if t == nil {
		return fmt.Errorf("bad index")
	}
	if t.Prev != nil {
		t.Prev.Next = t.Next
	} else {
		l.Head = t.Next
	}
	if t.Next != nil {
		t.Next.Prev = t.Prev
	} else {
		l.Tail = t.Prev
	}
	return nil
}

func (l *DoublyList) Find(val int) bool {
	for t := l.Head; t != nil; t = t.Next {
		if t.Val == val {
			return true
		}
	}
	return false
}

func (l *DoublyList) Serialize() string {
	var sb strings.Builder
	first := true
	for t := l.Head; t != nil; t = t.Next {
		if !first {
			sb.WriteString(" ")
		}
		sb.WriteString(strconv.Itoa(t.Val))
		first = false
	}
	return sb.String()
}

func (l *DoublyList) Deserialize(s string) {
	l.Head = nil
	l.Tail = nil
	parts := strings.Fields(s)
	for _, p := range parts {
		if v, err := strconv.Atoi(p); err == nil {
			l.PushTail(v)
		}
	}
}

func (l *DoublyList) ToString() string {
	var sb strings.Builder
	sb.WriteString("[ ")
	for t := l.Head; t != nil; t = t.Next {
		sb.WriteString(strconv.Itoa(t.Val) + " ")
	}
	sb.WriteString("]")
	return sb.String()
}
