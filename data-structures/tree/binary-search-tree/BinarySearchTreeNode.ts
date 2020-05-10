import BinaryTreeNode from '../binary-tree/BinaryTreeNode.ts'

class BinarySearchTreeNode<T> extends BinaryTreeNode<T> {
    constructor(value: T, left?: BinarySearchTreeNode<T>, right?: BinarySearchTreeNode<T>) {
        super(value, left, right)
    }

    find(value: T): BinarySearchTreeNode<T> | undefined {
        if (value === this.value) { // check the node itself
            return this
        }

        if (value < this.value && this.left instanceof BinarySearchTreeNode) { // check the left subtree
            return this.left?.find(value)
        }

        if (value > this.value && this.right instanceof BinarySearchTreeNode) { // check the right subtree
            return this.right?.find(value)
        }

        return undefined
    }
}

export default BinarySearchTreeNode
