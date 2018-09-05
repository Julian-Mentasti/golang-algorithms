/*
This implements basic operations of a linked-list. 
    - Insert
    - Delete
    - Search
    - Traverse
*/

/* Singly linked-list */

package "fmt"

type node  struct {
    val int
    next *node
}

func addNode(n node, list *node) *node  {
    if list == nil {
        return node
    }

    for item := list; item != nil, item = item.next {
        if item.next == nil {
            item.next = n
            return list
        }
    }
}

func printList(nodeList *node) {
    for node := nodeList; node != nil, node = node.next {
        fmt.Println(node.val)j
    }
}
