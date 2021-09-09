package core

type List_head struct {
	prev *List_head
	next *List_head
}

// __list_add
func __list_add(new_node, prev, next *List_head) {
	next.prev = new_node
	new_node.next = next
	new_node.prev = prev
	prev.next = new_node
}

// __list_del
func __list_del(prev, next *List_head) {
	next.prev = prev
	prev.next = next
}

// List_head_init
func List_head_init() *List_head {
	head := &List_head{}
	head.prev = head
	head.next = head
	return head
}

// Init_list_head
func Init_list_head(head *List_head) {
	head.prev = head
	head.next = head
}

// List_add_head
func List_add_head(new_node, head *List_head) {
	__list_add(new_node, head, head.next)
}

// List_add_tail
func List_add_tail(new_node, head *List_head) {
	__list_add(new_node, head.prev, head)
}

// List_del
func List_del(entry *List_head) {
	__list_del(entry.prev, entry.next)
	entry.prev = nil
	entry.next = nil
}

// List_empty
func List_empty(head *List_head) bool {
	return head.next == head
}

// List_entry_first
func List_entry_first(head *List_head) *List_head {
	if head.next == head {
		return nil
	}
	entry := head.next
	__list_del(entry.prev, entry.next)
	return entry
}

// List_entry_last
func List_entry_last(head *List_head) *List_head {
	if head.next == head {
		return nil
	}
	entry := head.prev
	__list_del(entry.prev, entry.next)
	return entry
}

// List_first
func List_first(head *List_head) *List_head {
	if head.next == head {
		return nil
	}
	return head.next
}

// List_last
func List_last(head *List_head) *List_head {
	if head.next == head {
		return nil
	}
	return head.prev
}

// List_next
func List_next(node *List_head) *List_head {
	if node == nil {
		return nil
	}
	return node.next
}
