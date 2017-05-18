package binaryTree

import(
    "testing"
)

func TestBinaryTree(t *testing.T) {
    bTree, err := NewLeaf(nil, nil, nil)
    if err != nil {
        t.Errorf("Recieved unexpected error from NewLeaf(): \n%s", err)
    }

    bTree.Insert(1)
    bTree.Insert(2)
    bTree.Insert(3)
    bTree.Insert(4)
    bTree.Insert(5)

    testVal := bTree.rightLeaf.rightLeaf.value
    if testVal != 3 {
        t.Errorf(
            "Failed to recieve the expected value of 3. Recieved: %d", testVal,
        )
    }

    err = bTree.Insert(50)
    err = bTree.Insert(510)
    err = bTree.Insert(20)
    err = bTree.Insert(19)
    err = bTree.Insert(33)

    if err != nil {
        t.Errorf("Recieved an unexpected error from Insert: %s", err)
    }

    desiredTree, err := bTree.Search(20)
    if err != nil {
        t.Errorf("Recieved an unexpected error from Search: %s", err)
    }

    if desiredTree.value != 20 {
        t.Errorf("Failed to find the correct value of 20. Recieved: %v", desiredTree.value)
    }

    emptyTree, err := bTree.Search(6)
    if err.Error() != "Unable to locate desired value : 6" {
        t.Error("Failed to recieve the expected error from Search")
    }

    if emptyTree != nil {
        t.Error("Failed to find the correct value of nil")
    }
}
