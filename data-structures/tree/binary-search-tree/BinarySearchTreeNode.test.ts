import BinarySearchTreeNode from './BinarySearchTreeNode.ts'

function makeTestSearchBinaryTree(): BinarySearchTreeNode<number> {
    const two = new BinarySearchTreeNode(2)
    const six = new BinarySearchTreeNode(6)
    const four = new BinarySearchTreeNode(4, two, six)
    const twenty = new BinarySearchTreeNode(20)
    const ten = new BinarySearchTreeNode(10, undefined, twenty)
    const eight = new BinarySearchTreeNode(8, four, ten)
    return eight
}

function testBinarySearchTreeFind() {
    const root = makeTestSearchBinaryTree()

    let search = root.find(6)
    console.assert(!!search, 'expect BinarySearchTree<>.find(6) to present')

    search = root.find(5)
    console.assert(!search, 'expect BinarySearchTree<>.find(5) to be blank')

    search = root.find(4)
    console.assert(!!search, 'expect BinarySearchTree<>.find(4) to present')

    search = root.find(20)
    console.assert(!!search, 'expect BinarySearchTree<>.find(20) to present')
}

testBinarySearchTreeFind()
