package singlylinkedlist

type Node struct {
	value int
	next  *Node
}
type LinkedList struct {
	head *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Get(index int) int {
	current := ll.head
	for i := 0; i < index; i++ {
		if current == nil {
			return -1
		}
		current = current.next
	}
	if current == nil {
		return -1
	}
	return current.value

}

func (ll *LinkedList) InsertHead(val int) {
	newNode := &Node{value: val, next: ll.head}
	if ll.head != nil {

		ll.head = newNode
	} else {
		ll.head = newNode
	}
	ll.size++

}

func (ll *LinkedList) InsertTail(val int) {
	current := ll.head
	for current != nil && current.next != nil {

		current = current.next
	}
	newNode := &Node{value: val}
	if current != nil {
		current.next = newNode
	} else {
		ll.head = newNode
	}
	ll.size++
}

func (ll *LinkedList) Remove(index int) bool {
	if index > ll.size {
		return false
	}
	current := ll.head
	i := 0
	for i = 0; i < index-1; i++ {
		if current == nil {
			return false
		}
		current = current.next
	}
	if current == nil {
		return false
	}
	if current.next == nil && i == index {
		ll.head = nil
		ll.size--
		return true
	}
	current.next = current.next.next
	ll.size--
	return true
}

func (ll *LinkedList) GetValues() []int {
	values := []int{}
	current := ll.head
	for current != nil {
		values = append(values, current.value)
		current = current.next
	}
	return values
}
