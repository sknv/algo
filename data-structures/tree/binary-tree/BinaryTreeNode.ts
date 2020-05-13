type Traverser<T> = (value: T) => void

class BinaryTreeNode<T> {
    constructor(public value: T, public left?: BinaryTreeNode<T>, public right?: BinaryTreeNode<T>) { }

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
