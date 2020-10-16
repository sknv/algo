package binary_tree

import (
	"reflect"
	"testing"
)

func TestBinaryTree_Invert(t *testing.T) {
	tree := makeBinaryTree()
	tree.Invert()

	var values []int
	tree.TraversePreOrder(func(value interface{}) {
		values = append(values, value.(int))
	})

	expected := []int{0, 2, 6, 1, 4, 3}
	if !reflect.DeepEqual(expected, values) {
		t.Errorf("expect inverted values to be [%v], got instead [%v]", expected, values)
	}
}

func makeBinaryTree() *BinaryTree {
	zero := BinaryTreeNode{
		Value: 0,
	}
	one := BinaryTreeNode{
		Value: 1,
	}
	two := BinaryTreeNode{
		Value: 2,
	}
	three := BinaryTreeNode{
		Value: 3,
	}
	four := BinaryTreeNode{
		Value: 4,
	}
	six := BinaryTreeNode{
		Value: 6,
	}

	zero.Left, zero.Right = &one, &two
	one.Left, one.Right = &three, &four
	two.Right = &six
	return &BinaryTree{
		Root: &zero,
	}
}
