package linkedList

import(
    "errors"
    "fmt"
    "sync"
)

type LinkedList struct {
    head    *Node
    last    *Node
    length  int
    locker  sync.RWMutex
}

type Node struct {
    Value   interface{}
    next    *Node
    prev    *Node
    list    *LinkedList
}

func NewLinkedList() (*LinkedList, error) {
    return &LinkedList{
        length: 0,
    }, nil
}

func (l *LinkedList) Insert(val interface{}) error {
    newNode := &Node{
        Value:  val,
        next:   l.head,
        prev:   l.last,
        list:   l,
    }

    l.locker.Lock()
    defer l.locker.Unlock()

    if l.head == nil {
        l.head = newNode
        l.last = newNode
    } else {
        l.head.prev = newNode
        l.head = newNode
        l.last.next = newNode
    }

    l.length++
    return nil
}

func (l *LinkedList) Find(val interface{}) (*Node, error) {
    defaultErr := errors.New(fmt.Sprintf("Unable to locate desired value: %v", val))
    if l.head == nil {
        return nil, defaultErr
    }

    currentNode := l.head

    for i := 1; i <= l.length; i++{
        if currentNode.Value == val {
            return currentNode, nil
        }
        if currentNode.next != nil {
            currentNode = currentNode.next
        }
    }
    return nil, defaultErr
}
