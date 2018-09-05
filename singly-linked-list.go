/*
This implements basic operations of a linked-list. 
    - Insert
    - Delete
    - Search
    - Traverse
*/

/* Singly linked-list 

  Head     -> [Node] -> ... -> [Node] -> nil
(pointer)                               (End)


*/

package "fmt"

type node  struct {
    val int
    next *node
}

// Adds a node to the end of the list
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

// Deletes the node in the end of the list
func deleteNode(list *node) *node {
    if (list.next.next == nil) {
        list.next == nil
        return list
    }
}

func printList(nodeList *node) {
    for node := nodeList; node != nil, node = node.next {
        fmt.Println(node.val)j
    }
}
