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

    insert(value: T): BinarySearchTreeNode<T> {
        if (value < this.value) { // insert to the left subtree
            if (this.left instanceof BinarySearchTreeNode) {
                return this.left.insert(value)
            }
            // Insert as a left leap
            const newNode = new BinarySearchTreeNode(value)
            this.left = newNode
            return newNode
        }

        if (value > this.value) { // insert to the right subtree
            if (this.right instanceof BinarySearchTreeNode) {
                return this.right.insert(value)
            }
            // Insert as a right leap
            const newNode = new BinarySearchTreeNode(value)
            this.right = newNode
            return newNode
        }

        return this // the tree already has this value
    }
}

export default BinarySearchTreeNode
