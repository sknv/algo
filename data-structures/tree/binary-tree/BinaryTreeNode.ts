type Traverser<T> = (value: T) => void

class BinaryTreeNode<T> {
    value: T
    left: BinaryTreeNode<T> | undefined
    right: BinaryTreeNode<T> | undefined

    constructor(value: T, left?: BinaryTreeNode<T>, right?: BinaryTreeNode<T>) {
        this.value = value
        this.left = left
        this.right = right
    }

    get leftHeight(): number {
        if (!this.left) {
            return 0
        }
        return this.left.height + 1
    }

    get rightHeight(): number {
        if (!this.right) {
            return 0
        }
        return this.right.height + 1
    }

    get height(): number {
        return Math.max(this.leftHeight, this.rightHeight)
    }

    get balanceFactor(): number {
        return this.leftHeight - this.rightHeight
    }

    traversePreOrder(traverse: Traverser<T>) {
        traverse(this.value)
        this.left?.traversePreOrder(traverse)
        this.right?.traversePreOrder(traverse)
    }

    traverseInOrder(traverse: Traverser<T>) {
        this.left?.traverseInOrder(traverse)
        traverse(this.value)
        this.right?.traverseInOrder(traverse)
    }

    traversePostOrder(traverse: Traverser<T>) {
        this.left?.traversePostOrder(traverse)
        this.right?.traversePostOrder(traverse)
        traverse(this.value)
    }
}

export default BinaryTreeNode
