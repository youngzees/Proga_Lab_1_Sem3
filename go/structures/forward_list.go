package structures

import (
	"fmt"
	"strconv"
	"strings"
)

type NodeF struct {
	Val  int
	Next *NodeF
}

type ForwardList struct {
	Head *NodeF
}

func (l *ForwardList) PushHead(val int) {
	l.Head = &NodeF{Val: val, Next: l.Head}
}

func (l *ForwardList) PushTail(val int) {
	n := &NodeF{Val: val}
	if l.Head == nil {
		l.Head = n
		return
	}
	t := l.Head
	for t.Next != nil {
		t = t.Next
	}
	t.Next = n
}

func (l *ForwardList) PushBefore(target, val int) error {
	if l.Head == nil {
		return fmt.Errorf("empty")
	}
	if l.Head.Val == target {
		l.PushHead(val)
		return nil
	}
	t := l.Head
	for t.Next != nil && t.Next.Val != target {
		t = t.Next
	}
	if t.Next == nil {
		return fmt.Errorf("target not found")
	}
	t.Next = &NodeF{Val: val, Next: t.Next}
	return nil
}

func (l *ForwardList) PushAfter(target, val int) error {
	t := l.Head
	for t != nil && t.Val != target {
		t = t.Next
	}
	if t == nil {
		return fmt.Errorf("target not found")
	}
	t.Next = &NodeF{Val: val, Next: t.Next}
	return nil
}

func (l *ForwardList) PopHead() error {
	if l.Head == nil {
		return fmt.Errorf("empty")
	}
	l.Head = l.Head.Next
	return nil
}

func (l *ForwardList) PopTail() error {
	if l.Head == nil {
		return fmt.Errorf("empty")
	}
	if l.Head.Next == nil {
		l.Head = nil
		return nil
	}
	t := l.Head
	for t.Next.Next != nil {
		t = t.Next
	}
	t.Next = nil
	return nil
}

func (l *ForwardList) Del(val int) error {
	if l.Head == nil {
		return fmt.Errorf("empty")
	}
	if l.Head.Val == val {
		l.Head = l.Head.Next
		return nil
	}
	t := l.Head
	for t.Next != nil && t.Next.Val != val {
		t = t.Next
	}
	if t.Next == nil {
		return fmt.Errorf("not found")
	}
	t.Next = t.Next.Next
	return nil
}

func (l *ForwardList) DelIndex(index int) error {
	if index < 0 || l.Head == nil {
		return fmt.Errorf("bad index")
	}
	if index == 0 {
		l.Head = l.Head.Next
		return nil
	}
	t := l.Head
	for i := 0; i < index-1 && t.Next != nil; i++ {
		t = t.Next
	}
	if t.Next == nil {
		return fmt.Errorf("bad index")
	}
	t.Next = t.Next.Next
	return nil
}

func (l *ForwardList) Find(val int) bool {
	for t := l.Head; t != nil; t = t.Next {
		if t.Val == val {
			return true
		}
	}
	return false
}

func (l *ForwardList) Serialize() string {
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

func (l *ForwardList) Deserialize(s string) {
	l.Head = nil
	parts := strings.Fields(s)
	for _, p := range parts {
		if v, err := strconv.Atoi(p); err == nil {
			l.PushTail(v)
		}
	}
}

func (l *ForwardList) ToString() string {
	var sb strings.Builder
	sb.WriteString("[ ")
	for t := l.Head; t != nil; t = t.Next {
		sb.WriteString(strconv.Itoa(t.Val) + " ")
	}
	sb.WriteString("]")
	return sb.String()
}
