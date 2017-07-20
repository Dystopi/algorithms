package linkedList

import(
    "testing"
)

func TestNewLinkedList(t *testing.T) {
    _, err := NewLinkedList()
    if err != nil {
        t.Errorf("Recieved an unexpected error from NewLinkedList: %s", err)
    }
}

func TestInsert(t *testing.T) {
    list, _ := NewLinkedList()
    list.Insert(1)

    if list.head.Value != 1 {
        t.Error("Failed to recieve the expected value at Head")
    }

    list.Insert(2)
    if list.last.Value != 1 {
        t.Error("Failed to recieve the expected value at Last")
    }

    if list.head.Value != 2 {
        t.Error("Failed to recieve the expected value at Head")
    }

}

func TestFind(t *testing.T) {
    list, _ := NewLinkedList()
    list.Insert(100)
    list.Insert(110)
    list.Insert(120)
    list.Insert(130)
    list.Insert(140)
    list.Insert(150)
    list.Insert(160)
    list.Insert(170)

    node, err := list.Find(3)
    if err.Error() != "Unable to locate desired value: 3" {
        t.Error("Failed to recieve the expected error when value not present")
    }

    node, err = list.Find(140)
    if err != nil {
        t.Errorf("Recieved an unexpected error from find: %s", err)
    }
    if node.Value != 140 {
        t.Error("Failed to recieved the expected value")
    }

    if node.next.Value != 130 {
        t.Error("Failed to recieved the expected value from the previous node")
    }

    if node.prev.Value != 150 {
        t.Error("Failed to recieved the expected value from the next node")
    }
}
