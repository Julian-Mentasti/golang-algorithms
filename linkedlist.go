/*
This implements the following operation of a linkedlist
 - Append
 - Insert
 - Delete
 - Remove
 - Search
 - Traverse

Singly linkedlist

    Head      -> [Node] -> ... -> [Node] -> nil
 (pointer)                                 (End) 
*/

package linkedlist

import "fmt"

//Node struct with a pointer to another node and an interger value
type Node struct {
     value int
     next  *node
}

//LinkedList struct with a node as head and a size value
type LinkedList struct {
    head *Node
    size int

}

//Append a node with the given value to the linked list
func (list *LinkedList) Append (value int) {
    newNode := Node{value, nil}

    if list.head == nil {
        list.head = &newNode
        return
    } else {
        temp := list.head
        for {
            if temp.head == nil {
                break
            }
            temp = temp.next
        }
        temp.next = &newNode
    }
    list.size++
}

//Insert a node at poition i.
func (list *ItemLinkedList) Insert(i int, value int) error {
    if i < 0 || i > list.size {
        return nil, fmt.Errorf("Out of bounds!")
    }
    node := list.head
    j := 0
    for j < i-1 { //Index out of 0
        j++
        node = node.next
    }
