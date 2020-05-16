import BinarySearchTreeNode from './BinarySearchTreeNode.ts'
import BinaryTree from '../binary-tree/BinaryTree.ts'
import BinaryTreeNode from '../binary-tree/BinaryTreeNode.ts'

class BinarySearchTree<T> extends BinaryTree<T>{
    constructor(rootValue?: T) {
        let root: BinaryTreeNode<T> | undefined
        if (rootValue) {
            root = new BinaryTreeNode(rootValue)
        }
        super(root)
    }

    find(value: T): BinaryTreeNode<T> | undefined {
        return BinarySearchTreeNode.find(value, this.root)
    }

    insert(value: T): BinaryTreeNode<T> {
        if (!this.root) { // if there is no root yet
            this.root = BinarySearchTreeNode.insert(value)
            return this.root
        }

        return BinarySearchTreeNode.insert(value, this.root)
    }

    delete(value: T): BinaryTreeNode<T> | undefined {
        return BinarySearchTreeNode.delete(value, this.root)
    }
}

export default BinarySearchTree
