package main

import "fmt"

type Node struct {
    data int
    next *Node
}

type LinkedList struct {
    head *Node
}

func (ll *LinkedList) insert(data int) {
    newNode := &Node{
        data: data,
        next: nil,
    }

    if ll.head == nil {
        ll.head = newNode
    } else {
        current := ll.head
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
    }
}

func (ll *LinkedList) display() {
    if ll.head == nil {
        fmt.Println("Linked list is empty")
        return
    }

    current := ll.head
    for current != nil {
        fmt.Printf("%d ", current.data)
        current = current.next
    }
    fmt.Println()
}

func main() {
    l := LinkedList{}
    l.insert(10)
    l.display()


}

