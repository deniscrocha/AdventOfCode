package main

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	length int
}

func (list *LinkedList) insertAtBack(data int) {
	newNode := &Node{data: data, next: nil}
	if list.head == nil {
		list.head = newNode
		list.length += 1
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	list.length += 1
}

func (list *LinkedList) insertAtFront(data int) {
	if list.head == nil {
		newNode := &Node{data: data, next: nil}
		list.length += 1
		list.head = newNode
		return
	}
	newNode := &Node{data: data, next: list.head}
	list.length += 1
	list.head = newNode
}

func (list *LinkedList) get(index int) (int){
	if list.head == nil{
		panic("Index not found")
	}
	if index > list.length{
		panic("Index out of bound")
	}
	counter := 0
	current := list.head
	for current.next != nil {
		if counter == index{
			return current.data
		}
		counter++
		current = current.next
	}
	return current.data
}
func (list *LinkedList) set(index int, data int){
	if list.head == nil{
		panic("Index not found")
	}
	if index > list.length{
		panic("Index out of bound")
	}
	counter := 0
	current := list.head
	for current.next != nil {
		if counter == index{
			current.data = data
			return
		}
		counter++
		current = current.next
	}
	current.data = data
}

func (list *LinkedList) bubbleSort(){
	for i:=0; i< list.length-1; i++{
		for j:=0; j < list.length-i-1; j++{
			if(list.get(j) > list.get(j+1)){
				lowerValue := list.get(j+1)
				value := list.get(j)
				list.set(j, lowerValue)
				list.set(j+1, value)
			}
		}
	}
}
