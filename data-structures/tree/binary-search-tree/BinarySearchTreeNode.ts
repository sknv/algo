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

    findMin(): BinarySearchTreeNode<T> {
        let current: BinarySearchTreeNode<T> = this
        while (current.left instanceof BinarySearchTreeNode) {
            current = current.left
        }
        return current
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

    delete(value: T): BinarySearchTreeNode<T> | undefined {
        if (value < this.value && this.left instanceof BinarySearchTreeNode) { // value is in the left subtree
            const childNode = this.left.delete(value)
            if (!childNode) { return undefined }
            this.left = childNode
            return this
        }

        if (value > this.value && this.right instanceof BinarySearchTreeNode) { // value is in the right subtree
            const childNode = this.right.delete(value)
            if (!childNode) { return undefined }
            this.right = childNode
            return this
        }

        if (value !== this.value) {
            return undefined
        }

        // Node with only one child or no child
        if (!this.left) {
            return this.right as BinarySearchTreeNode<T>
        }
        if (!this.right) {
            return this.left as BinarySearchTreeNode<T>
        }

        // Node has both children
        const minNode = (this.right as BinarySearchTreeNode<T>).findMin() // get the smallest in the right subtree
        this.value = minNode.value // copy the inorder successor's content to this node
        this.right = (this.right as BinarySearchTreeNode<T>).delete(minNode.value) // delete the inorder successor
        return this
    }
}

export default BinarySearchTreeNode
