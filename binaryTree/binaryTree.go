package binaryTree

import(
    "errors"
    "fmt"
    "reflect"
)

type lessCallback    func(b *BinaryTree, new interface{}) (bool, error)
type insertCallback  func(b *BinaryTree, new interface{}) error
type searchCallback  func(b *BinaryTree, desired interface{}) (*BinaryTree, error)

type BinaryTree struct {
    value           interface{}
    leftLeaf        *BinaryTree
    rightLeaf       *BinaryTree
    insert          insertCallback
    search          searchCallback
    less            lessCallback
}

func NewLeaf(comp lessCallback, ins insertCallback, srch searchCallback) (*BinaryTree, error) {
    if comp == nil {
        comp = defaultComparisonCallback
    }
    if ins == nil {
        ins = defaultInsert
    }
    if srch == nil {
        srch = defaultSearch
    }
    return &BinaryTree{
        value:          nil,
        less:           comp,
        insert:         ins,
        search:         srch,
    }, nil
}

// Not attaching logic directly to struct to allow for ease of testing
// new Insertion / Search / Comparison algorithms
func (b *BinaryTree) Insert(newValue interface{}) error {
    return b.insert(b, newValue)
}

func (b *BinaryTree) Search(desired interface{}) (*BinaryTree, error) {
    return b.search(b, desired)
}

func defaultInsert(b *BinaryTree, newValue interface{}) error {
    if b.value == nil {
        b.value = newValue
        b.rightLeaf, _  = NewLeaf(b.less, b.insert, b.search)
        b.leftLeaf, _   = NewLeaf(b.less, b.insert, b.search)
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

func defaultSearch(b *BinaryTree, desired interface{}) (*BinaryTree, error) {
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
