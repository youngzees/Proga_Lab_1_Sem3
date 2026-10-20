package structures

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FullBinaryTree struct {
	Root *TreeNode
}

func (t *FullBinaryTree) Insert(val int) {
	n := &TreeNode{Val: val}
	if t.Root == nil {
		t.Root = n
		return
	}
	queue := []*TreeNode{t.Root}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.Left == nil {
			cur.Left = n
			return
		}
		queue = append(queue, cur.Left)
		if cur.Right == nil {
			cur.Right = n
			return
		}
		queue = append(queue, cur.Right)
	}
}

func (t *FullBinaryTree) Find(val int) bool {
	if t.Root == nil {
		return false
	}
	queue := []*TreeNode{t.Root}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.Val == val {
			return true
		}
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	return false
}

func (t *FullBinaryTree) IsFull() bool {
	return checkFull(t.Root)
}

func checkFull(n *TreeNode) bool {
	if n == nil {
		return true
	}
	if n.Left == nil && n.Right == nil {
		return true
	}
	if n.Left != nil && n.Right != nil {
		return checkFull(n.Left) && checkFull(n.Right)
	}
	return false
}

func serTree(n *TreeNode) string {
	if n == nil {
		return "-"
	}
	if n.Left == nil && n.Right == nil {
		return strconv.Itoa(n.Val)
	}
	return fmt.Sprintf("%d(%s)(%s)", n.Val, serTree(n.Left), serTree(n.Right))
}

func (t *FullBinaryTree) Serialize() string {
	return serTree(t.Root)
}

func parseNode(s string, pos int) (*TreeNode, int, error) {
	if pos >= len(s) {
		return nil, pos, fmt.Errorf("bad tree format")
	}
	if s[pos] == '-' {
		return nil, pos + 1, nil
	}

	val := 0
	has := false
	for pos < len(s) && s[pos] >= '0' && s[pos] <= '9' {
		val = val*10 + int(s[pos]-'0')
		pos++
		has = true
	}
	if !has {
		return nil, pos, fmt.Errorf("bad tree format: no number")
	}

	node := &TreeNode{Val: val}
	if pos < len(s) && s[pos] == '(' {
		pos++
		var err error
		node.Left, pos, err = parseNode(s, pos)
		if err != nil {
			return nil, pos, err
		}
		if pos >= len(s) || s[pos] != ')' {
			return nil, pos, fmt.Errorf("expected )")
		}
		pos++
		if pos >= len(s) || s[pos] != '(' {
			return nil, pos, fmt.Errorf("expected (")
		}
		pos++
		node.Right, pos, err = parseNode(s, pos)
		if err != nil {
			return nil, pos, err
		}
		if pos >= len(s) || s[pos] != ')' {
			return nil, pos, fmt.Errorf("expected )")
		}
		pos++
	}
	return node, pos, nil
}

func (t *FullBinaryTree) Deserialize(s string) error {
	t.Root = nil
	if s == "" {
		return nil
	}
	var err error
	t.Root, _, err = parseNode(s, 0)
	return err
}

func (t *FullBinaryTree) ToString() string {
	return t.Serialize()
}
