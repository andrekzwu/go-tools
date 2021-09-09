package core

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type ListDemo struct {
	Value int
	List_head
}

func TestListHead(t *testing.T) {
	head := List_head_init()
	// n1
	n1 := &ListDemo{Value: 1, List_head: List_head{}}
	List_add_tail(&n1.List_head, head)
	// n2
	n2 := &ListDemo{Value: 2, List_head: List_head{}}
	List_add_tail(&n2.List_head, head)
	// n3
	n3 := &ListDemo{Value: 3, List_head: List_head{}}
	List_add_tail(&n3.List_head, head)
	// print
	PrintListDemo(head)
	return
}

// Listhead2ListDemo
func Listhead2ListDemo(pnode *List_head, rt reflect.Type) *ListDemo {
	f, ok := rt.FieldByName("List_head")
	if !ok {
		return nil
	}
	return (*ListDemo)(unsafe.Pointer(uintptr(unsafe.Pointer(pnode)) - f.Offset))
}

// PrintListDemo
func PrintListDemo(head *List_head) {
	h := List_next(head)
	for h != head {
		n := Listhead2ListDemo(h, reflect.TypeOf((*ListDemo)(nil)).Elem())
		fmt.Println(n.Value)
		h = List_next(h)
	}
}
