import BinaryTreeNode from './BinaryTreeNode.ts'

function makeTestBinaryTree(): BinaryTreeNode<number> {
    const two = new BinaryTreeNode(2)
    const six = new BinaryTreeNode(6)
    const four = new BinaryTreeNode(4, two, six)
    const ten = new BinaryTreeNode(10)
    const eight = new BinaryTreeNode(8, four, ten)
    return eight
}

function testBinaryTree() {
    const root = makeTestBinaryTree()

    const height = root.height
    let expected = 2
    console.assert(height === expected, `expect BinaryTree<>.height to be: ${expected}, got instead: ${height}`)

    const leftHeight = root.leftHeight
    console.assert(leftHeight === expected, `expect BinaryTree<>.leftHeight to be: ${expected}, got instead: ${leftHeight}`)

    const rightHeight = root.rightHeight
    expected = 1
    console.assert(rightHeight === expected, `expect BinaryTree<>.rightHeight to be: ${expected}, got instead: ${rightHeight}`)

    const balanceFactor = root.balanceFactor
    expected = 1
    console.assert(balanceFactor === expected, `expect BinaryTree<>.balanceFactor to be: ${expected}, got instead: ${balanceFactor}`)
}

function testBinaryTreeTraversePreOrder() {
    const root = makeTestBinaryTree()

    let nums: number[] = [], expected = [8, 4, 2, 6, 10]
    root.traversePreOrder((value: number) => nums.push(value))
    console.assert(JSON.stringify(nums) === JSON.stringify(expected), `expect BinaryTree<>.traversePreOrder() to be: ${expected}, got instead: ${nums}`)
}

function testBinaryTreeTraverseInOrder() {
    const root = makeTestBinaryTree()

    let nums: number[] = [], expected = [2, 4, 6, 8, 10]
    root.traverseInOrder((value: number) => nums.push(value))
    console.assert(JSON.stringify(nums) === JSON.stringify(expected), `expect BinaryTree<>.traverseInOrder() to be: ${expected}, got instead: ${nums}`)
}

function testBinaryTreeTraversePostOrder() {
    const root = makeTestBinaryTree()

    let nums: number[] = [], expected = [2, 6, 4, 10, 8]
    root.traversePostOrder((value: number) => nums.push(value))
    console.assert(JSON.stringify(nums) === JSON.stringify(expected), `expect BinaryTree<>.traversePostOrder() to be: ${expected}, got instead: ${nums}`)
}

testBinaryTree()
testBinaryTreeTraversePreOrder()
testBinaryTreeTraverseInOrder()
testBinaryTreeTraversePostOrder()
