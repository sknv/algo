import BinaryTreeNode from "../binary-tree/BinaryTreeNode.ts";

class BinarySearchTreeNode {
  static find<T>(
    value: T,
    node?: BinaryTreeNode<T>,
  ): BinaryTreeNode<T> | undefined {
    if (!node) {
      return undefined;
    }

    if (value === node.value) { // check the node itself
      return node;
    }

    if (value < node.value) { // check the left subtree
      return BinarySearchTreeNode.find(value, node.left);
    }

    if (value > node.value) { // check the right subtree
      return BinarySearchTreeNode.find(value, node.right);
    }

    return undefined;
  }

  static findMin<T>(node?: BinaryTreeNode<T>): BinaryTreeNode<T> | undefined {
    if (!node) {
      return undefined;
    }

    let current: BinaryTreeNode<T> = node;
    while (current.left) {
      current = current.left;
    }
    return current;
  }

  static insert<T>(value: T, node?: BinaryTreeNode<T>): BinaryTreeNode<T> {
    if (!node) {
      return new BinaryTreeNode(value);
    }

    if (value < node.value) { // insert to the left subtree
      if (node.left) {
        return BinarySearchTreeNode.insert(value, node.left);
      }
      // Insert as a left leap
      const newNode = new BinaryTreeNode(value);
      node.left = newNode;
      return newNode;
    }

    if (value > node.value) { // insert to the right subtree
      if (node.right) {
        return BinarySearchTreeNode.insert(value, node.right);
      }
      // Insert as a right leap
      const newNode = new BinaryTreeNode(value);
      node.right = newNode;
      return newNode;
    }

    return node; // the tree already has this value
  }

  static delete<T>(
    value: T,
    node?: BinaryTreeNode<T>,
  ): BinaryTreeNode<T> | undefined {
    if (!node) {
      return undefined;
    }

    if (value < node.value) { // value is in the left subtree
      const childNode = BinarySearchTreeNode.delete(value, node.left);
      if (!childNode) {
        return undefined;
      }
      node.left = childNode;
      return node;
    }

    if (value > node.value) { // value is in the right subtree
      const childNode = BinarySearchTreeNode.delete(value, node.right);
      if (!childNode) {
        return undefined;
      }
      node.right = childNode;
      return node;
    }

    if (value !== node.value) {
      return undefined;
    }

    // Node with only one child or no child
    if (!node.left) {
      return node.right;
    }
    if (!node.right) {
      return node.left;
    }

    // Node has both children
    const minNode = BinarySearchTreeNode.findMin(node.right); // get the smallest in the right subtree
    node.value = minNode!.value; // copy the inorder successor's content to this node
    node.right = BinarySearchTreeNode.delete(minNode!.value, node.right); // delete the inorder successor
    return node;
  }
}

export default BinarySearchTreeNode;
