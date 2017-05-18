package binaryTree

import(
    "errors"
    "fmt"
    "reflect"
)

type lessCallback    func(b *BinaryTree, new interface{}) (bool, error)

type BinaryTree struct {
    value           interface{}
    leftLeaf        *BinaryTree
    rightLeaf       *BinaryTree
    less            lessCallback
}

func NewLeaf(comp lessCallback) (*BinaryTree, error) {
    if comp == nil {
        comp = defaultComparisonCallback
    }
    return &BinaryTree{
        value:          nil,
        less:           comp,
    }, nil
}

func (b *BinaryTree) Insert(newValue interface{}) error {
    if b.value == nil {
        b.value = newValue
        b.rightLeaf, _  = NewLeaf(b.less)
        b.leftLeaf, _   = NewLeaf(b.less)
    } else {
        isLess, err := b.less(b, newValue)
        if  err != nil {
            return err
        }

        if isLess == true {
            b.leftLeaf.Insert(newValue)
        } else {
            b.rightLeaf.Insert(newValue)
        }
    }
    return nil
}

func (b *BinaryTree) Search(desired interface{}) (*BinaryTree, error) {
    if b.value == nil {
        return nil, errors.New(fmt.Sprintf("Unable to locate desired value : %v", desired))
    }

    if reflect.TypeOf(b.value).Name() != reflect.TypeOf(desired).Name() {
        return nil, errors.New("Divergent Interface Types")
    }

    if b.value == desired {
        return b, nil
    } else {
        isLess, err := b.less(b, desired)
        if err != nil {
            return nil, err
        }

        if isLess {
            return b.leftLeaf.Search(desired)
        }
    }
    return b.rightLeaf.Search(desired)
}

func defaultComparisonCallback(b *BinaryTree, newValue interface{}) (bool, error) {
    if reflect.TypeOf(newValue).Name() != reflect.TypeOf(b.value).Name() {
        return false, errors.New("Divergent Interface Types")
    }

    switch newValue.(type) {
        case int:
            if newValue.(int) < b.value.(int) {
                return true, nil
            } else {
                return false, nil
            }
        default:
            return false, errors.New("Unknown Interface Type")
    }
}
