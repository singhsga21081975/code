package main

import "fmt"

type Node struct {
    data int
    next *Node
}

type LL struct {
    head *Node
}

// Method with a pointer receiver: Modifies the linked list
func (l *LL) InsertPointer(data int) {
    newNode := &Node{data: data}

    if l.head == nil {
        l.head = newNode
    } else {
        current := l.head
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
    }
}

// Method with a value receiver: Does not modify the original linked list
func (l LL) InsertValue(data int) {
    newNode := &Node{data: data}

    if l.head == nil {
        l.head = newNode
    } else {
        current := l.head
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
    }
}

func main() {
    list1 := LL{} // Creating a linked list instance

    // Using the method with a pointer receiver
    list1.InsertPointer(10)
    list1.InsertPointer(20)
    fmt.Println("List1 after InsertPointer:", list1.head.data, list1.head.next.data) // Output: 10 20

    list2 := LL{} // Creating another linked list instance

    // Using the method with a value receiver
    list2.InsertValue(30)
    list2.InsertValue(40)
    //fmt.Println("List2 after InsertValue:", list2.head.data, list2.head.next.data) // Output: 30 40
}
