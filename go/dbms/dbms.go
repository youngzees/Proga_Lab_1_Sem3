package dbms

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"dbms/structures"
)

type DBMS struct {
	Filename string

	Arrays map[string]*structures.MyArray
	FLists map[string]*structures.ForwardList
	DLists map[string]*structures.DoublyList
	Stacks map[string]*structures.MyStack
	Queues map[string]*structures.MyQueue
	Trees  map[string]*structures.FullBinaryTree
}

func NewDBMS(filename string) *DBMS {
	db := &DBMS{
		Filename: filename,
		Arrays:   make(map[string]*structures.MyArray),
		FLists:   make(map[string]*structures.ForwardList),
		DLists:   make(map[string]*structures.DoublyList),
		Stacks:   make(map[string]*structures.MyStack),
		Queues:   make(map[string]*structures.MyQueue),
		Trees:    make(map[string]*structures.FullBinaryTree),
	}
	db.load()
	return db
}

func (db *DBMS) getArray(name string) *structures.MyArray {
	if _, ok := db.Arrays[name]; !ok {
		db.Arrays[name] = &structures.MyArray{}
	}
	return db.Arrays[name]
}

func (db *DBMS) getFList(name string) *structures.ForwardList {
	if _, ok := db.FLists[name]; !ok {
		db.FLists[name] = &structures.ForwardList{}
	}
	return db.FLists[name]
}

func (db *DBMS) getDList(name string) *structures.DoublyList {
	if _, ok := db.DLists[name]; !ok {
		db.DLists[name] = &structures.DoublyList{}
	}
	return db.DLists[name]
}

func (db *DBMS) getStack(name string) *structures.MyStack {
	if _, ok := db.Stacks[name]; !ok {
		db.Stacks[name] = &structures.MyStack{}
	}
	return db.Stacks[name]
}

func (db *DBMS) getQueue(name string) *structures.MyQueue {
	if _, ok := db.Queues[name]; !ok {
		db.Queues[name] = &structures.MyQueue{}
	}
	return db.Queues[name]
}

func (db *DBMS) getTree(name string) *structures.FullBinaryTree {
	if _, ok := db.Trees[name]; !ok {
		db.Trees[name] = &structures.FullBinaryTree{}
	}
	return db.Trees[name]
}

func (db *DBMS) load() {
	file, err := os.Open(db.Filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, " ", 3)
		if len(parts) < 2 {
			continue
		}
		tag, name := parts[0], parts[1]
		data := ""
		if len(parts) == 3 {
			data = parts[2]
		}

		switch tag {
		case "M":
			db.getArray(name).Deserialize(data)
		case "F":
			db.getFList(name).Deserialize(data)
		case "L":
			db.getDList(name).Deserialize(data)
		case "S":
			db.getStack(name).Deserialize(data)
		case "Q":
			db.getQueue(name).Deserialize(data)
		case "T":
			db.getTree(name).Deserialize(data)
		}
	}
}

func (db *DBMS) Save() {
	file, err := os.Create(db.Filename)
	if err != nil {
		return
	}
	defer file.Close()

	for name, a := range db.Arrays {
		fmt.Fprintf(file, "M %s %s\n", name, a.Serialize())
	}
	for name, l := range db.FLists {
		fmt.Fprintf(file, "F %s %s\n", name, l.Serialize())
	}
	for name, l := range db.DLists {
		fmt.Fprintf(file, "L %s %s\n", name, l.Serialize())
	}
	for name, s := range db.Stacks {
		fmt.Fprintf(file, "S %s %s\n", name, s.Serialize())
	}
	for name, q := range db.Queues {
		fmt.Fprintf(file, "Q %s %s\n", name, q.Serialize())
	}
	for name, t := range db.Trees {
		fmt.Fprintf(file, "T %s %s\n", name, t.Serialize())
	}
}

func (db *DBMS) printAll() {
	any := false
	for name, a := range db.Arrays {
		fmt.Printf("ARRAY    %s : %s\n", name, a.Serialize())
		any = true
	}
	for name, l := range db.FLists {
		fmt.Printf("FLIST    %s : %s\n", name, l.ToString())
		any = true
	}
	for name, l := range db.DLists {
		fmt.Printf("DLIST    %s : %s\n", name, l.ToString())
		any = true
	}
	for name, s := range db.Stacks {
		fmt.Printf("STACK    %s : %s\n", name, s.ToString())
		any = true
	}
	for name, q := range db.Queues {
		fmt.Printf("QUEUE    %s : %s\n", name, q.ToString())
		any = true
	}
	for name, t := range db.Trees {
		fmt.Printf("TREE     %s : %s\n", name, t.ToString())
		any = true
	}
	if !any {
		fmt.Println("(empty)")
	}
}

func (db *DBMS) Execute(query string) {
	parts := strings.Fields(query)
	if len(parts) == 0 {
		return
	}
	cmd := parts[0]

	if cmd == "PRINT" {
		db.printAll()
		return
	}

	if len(parts) < 2 {
		fmt.Println("Error: no structure name")
		return
	}
	name := parts[1]

	// Хелпер для парсинга интов
	parseInt := func(idx int) (int, bool) {
		if idx >= len(parts) {
			return 0, false
		}
		v, err := strconv.Atoi(parts[idx])
		return v, err == nil
	}

	switch cmd {
	// ---- M ----
	case "MPUSH":
		if v, ok := parseInt(2); ok {
			db.getArray(name).Push(v)
		}
	case "MDEL":
		if v, ok := parseInt(2); ok {
			if err := db.getArray(name).Del(v); err != nil {
				fmt.Println("Error:", err)
			}
		}
	case "MGET":
		if v, ok := parseInt(2); ok {
			if res, err := db.getArray(name).Get(v); err == nil {
				fmt.Println(res)
			} else {
				fmt.Println("Error:", err)
			}
		}
	case "MSET":
		if i, ok1 := parseInt(2); ok1 {
			if v, ok2 := parseInt(3); ok2 {
				if err := db.getArray(name).Set(i, v); err != nil {
					fmt.Println("Error:", err)
				}
			}
		}
	case "MLEN":
		fmt.Println(db.getArray(name).Length())
	case "MPRINT":
		fmt.Println(db.getArray(name).Serialize())

	// ---- F ----
	case "FADDH":
		if v, ok := parseInt(2); ok {
			db.getFList(name).PushHead(v)
		}
	case "FADDT", "FPUSH":
		if v, ok := parseInt(2); ok {
			db.getFList(name).PushTail(v)
		}
	case "FADDB":
		if t, ok1 := parseInt(2); ok1 {
			if v, ok2 := parseInt(3); ok2 {
				if err := db.getFList(name).PushBefore(t, v); err != nil {
					fmt.Println("Error:", err)
				}
			}
		}
	case "FADDA":
		if t, ok1 := parseInt(2); ok1 {
			if v, ok2 := parseInt(3); ok2 {
				if err := db.getFList(name).PushAfter(t, v); err != nil {
					fmt.Println("Error:", err)
				}
			}
		}
	case "FDELH":
		if err := db.getFList(name).PopHead(); err != nil {
			fmt.Println("Error:", err)
		}
	case "FDELT":
		if err := db.getFList(name).PopTail(); err != nil {
			fmt.Println("Error:", err)
		}
	case "FDEL":
		if v, ok := parseInt(2); ok {
			if err := db.getFList(name).Del(v); err != nil {
				fmt.Println("Error:", err)
			}
		}
	case "FDELI":
		if v, ok := parseInt(2); ok {
			if err := db.getFList(name).DelIndex(v); err != nil {
				fmt.Println("Error:", err)
			}
		}
	case "FGET":
		if v, ok := parseInt(2); ok {
			fmt.Println(db.getFList(name).Find(v))
		}
	case "FPRINT":
		fmt.Println(db.getFList(name).ToString())

	// ---- L ----
	case "LADDH":
		if v, ok := parseInt(2); ok {
			db.getDList(name).PushHead(v)
		}
	case "LADDT", "LPUSH":
		if v, ok := parseInt(2); ok {
			db.getDList(name).PushTail(v)
		}
	case "LADDB":
		if t, ok1 := parseInt(2); ok1 {
			if v, ok2 := parseInt(3); ok2 {
				if err := db.getDList(name).PushBefore(t, v); err != nil {
					fmt.Println("Error:", err)
				}
			}
		}
	case "LADDA":
		if t, ok1 := parseInt(2); ok1 {
			if v, ok2 := parseInt(3); ok2 {
				if err := db.getDList(name).PushAfter(t, v); err != nil {
					fmt.Println("Error:", err)
				}
			}
		}
	case "LDELH":
		if err := db.getDList(name).PopHead(); err != nil {
			fmt.Println("Error:", err)
		}
	case "LDELT":
		if err := db.getDList(name).PopTail(); err != nil {
			fmt.Println("Error:", err)
		}
	case "LDEL":
		if v, ok := parseInt(2); ok {
			if err := db.getDList(name).Del(v); err != nil {
				fmt.Println("Error:", err)
			}
		}
	case "LDELI":
		if v, ok := parseInt(2); ok {
			if err := db.getDList(name).DelIndex(v); err != nil {
				fmt.Println("Error:", err)
			}
		}
	case "LGET":
		if v, ok := parseInt(2); ok {
			fmt.Println(db.getDList(name).Find(v))
		}
	case "LPRINT":
		fmt.Println(db.getDList(name).ToString())

	// ---- S ----
	case "SPUSH":
		if v, ok := parseInt(2); ok {
			db.getStack(name).Push(v)
		}
	case "SPOP":
		if v, err := db.getStack(name).Pop(); err == nil {
			fmt.Println(v)
		} else {
			fmt.Println("Error:", err)
		}
	case "SPRINT":
		fmt.Println(db.getStack(name).ToString())

	// ---- Q ----
	case "QPUSH":
		if v, ok := parseInt(2); ok {
			db.getQueue(name).Push(v)
		}
	case "QPOP":
		if v, err := db.getQueue(name).Pop(); err == nil {
			fmt.Println(v)
		} else {
			fmt.Println("Error:", err)
		}
	case "QPRINT":
		fmt.Println(db.getQueue(name).ToString())

	// ---- T ----
	case "TINSERT":
		if v, ok := parseInt(2); ok {
			db.getTree(name).Insert(v)
		}
	case "TGET":
		if v, ok := parseInt(2); ok {
			fmt.Println(db.getTree(name).Find(v))
		}
	case "TISFULL":
		fmt.Println(db.getTree(name).IsFull())
	case "TPRINT":
		fmt.Println(db.getTree(name).ToString())

	default:
		fmt.Println("Unknown command:", cmd)
	}
}
