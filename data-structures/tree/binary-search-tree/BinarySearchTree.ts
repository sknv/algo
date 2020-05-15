import BinarySearchTreeNode from './BinarySearchTreeNode.ts'

class BinarySearchTree<T> {
    root?: BinarySearchTreeNode<T>

    constructor(rootValue?: T) {
        if (rootValue) {
            this.root = new BinarySearchTreeNode(rootValue)
        }
    }

    find(value: T): BinarySearchTreeNode<T> | undefined {
        return this.root?.find(value)
    }

    insert(value: T): BinarySearchTreeNode<T> {
        if (!this.root) { // if there is no root yet
            this.root = new BinarySearchTreeNode(value)
            return this.root
        }

        return this.root.insert(value)
    }

    delete(value: T): BinarySearchTreeNode<T> | undefined {
        return this.root?.delete(value)
    }
}

export default BinarySearchTree
