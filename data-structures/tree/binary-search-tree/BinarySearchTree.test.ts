import BinarySearchTree from "./BinarySearchTree.ts";

function makeTestSearchBinaryTree(): BinarySearchTree<number> {
  const tree = new BinarySearchTree(8);
  tree.insert(10);
  tree.insert(20);
  tree.insert(6);
  tree.insert(4);
  tree.insert(2);
  return tree;
}

function testBinarySearchTreeInsert() {
  const tree = makeTestSearchBinaryTree();

  let nums: number[] = [], expected = [2, 4, 6, 8, 10, 20];
  tree.traverseInOrder((value: number) => nums.push(value));
  console.assert(
    JSON.stringify(nums) === JSON.stringify(expected),
    `expect BinarySearchTree<>.root.traverseInOrder() to be: ${expected}, got instead: ${nums}`,
  );
}

function testBinarySearchTreeFind() {
  const root = makeTestSearchBinaryTree();

  let search = root.find(6);
  console.assert(!!search, "expect BinarySearchTree<>.find(6) to present");

  search = root.find(5);
  console.assert(!search, "expect BinarySearchTree<>.find(5) to be blank");

  search = root.find(4);
  console.assert(!!search, "expect BinarySearchTree<>.find(4) to present");

  search = root.find(20);
  console.assert(!!search, "expect BinarySearchTree<>.find(20) to present");
}

function testBinarySearchTreeDelete() {
  const tree = makeTestSearchBinaryTree();

  tree.delete(4);
  let nums: number[] = [], expected = [2, 6, 8, 10, 20];
  tree.traverseInOrder((value: number) => nums.push(value));
  console.assert(
    JSON.stringify(nums) === JSON.stringify(expected),
    `expect BinarySearchTree<>.root.traverseInOrder() to be: ${expected}, got instead: ${nums}`,
  );

  tree.delete(9);
  nums = [], expected = [2, 6, 8, 10, 20];
  tree.traverseInOrder((value: number) => nums.push(value));
  console.assert(
    JSON.stringify(nums) === JSON.stringify(expected),
    `expect BinarySearchTree<>.root.traverseInOrder() to be: ${expected}, got instead: ${nums}`,
  );

  tree.delete(8);
  nums = [], expected = [2, 6, 10, 20];
  tree.traverseInOrder((value: number) => nums.push(value));
  console.assert(
    JSON.stringify(nums) === JSON.stringify(expected),
    `expect BinarySearchTree<>.root.traverseInOrder() to be: ${expected}, got instead: ${nums}`,
  );

  tree.delete(7);
  nums = [], expected = [2, 6, 10, 20];
  tree.traverseInOrder((value: number) => nums.push(value));
  console.assert(
    JSON.stringify(nums) === JSON.stringify(expected),
    `expect BinarySearchTree<>.root.traverseInOrder() to be: ${expected}, got instead: ${nums}`,
  );

  tree.delete(10);
  nums = [], expected = [2, 6, 20];
  tree.traverseInOrder((value: number) => nums.push(value));
  console.assert(
    JSON.stringify(nums) === JSON.stringify(expected),
    `expect BinarySearchTree<>.root.traverseInOrder() to be: ${expected}, got instead: ${nums}`,
  );
}

testBinarySearchTreeInsert();
testBinarySearchTreeFind();
testBinarySearchTreeDelete();
