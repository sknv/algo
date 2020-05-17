type Traverser<T> = (value: T) => void;

class BinaryTreeNode<T> {
  constructor(
    public value: T,
    public left?: BinaryTreeNode<T>,
    public right?: BinaryTreeNode<T>,
  ) {}

  static getLeftHeight<T>(node?: BinaryTreeNode<T>): number {
    if (!node?.left) {
      return 0;
    }
    return BinaryTreeNode.getHeight(node.left) + 1;
  }

  static getRightHeight<T>(node?: BinaryTreeNode<T>): number {
    if (!node?.right) {
      return 0;
    }
    return BinaryTreeNode.getHeight(node.right) + 1;
  }

  static getHeight<T>(node?: BinaryTreeNode<T>): number {
    return Math.max(
      BinaryTreeNode.getLeftHeight(node),
      BinaryTreeNode.getRightHeight(node),
    );
  }

  static getBalanceFactor<T>(node?: BinaryTreeNode<T>): number {
    return BinaryTreeNode.getLeftHeight(node) -
      BinaryTreeNode.getRightHeight(node);
  }

  static traversePreOrder<T>(traverse: Traverser<T>, node?: BinaryTreeNode<T>) {
    if (!node) {
      return;
    }

    traverse(node.value);
    BinaryTreeNode.traversePreOrder(traverse, node.left);
    BinaryTreeNode.traversePreOrder(traverse, node.right);
  }

  static traverseInOrder<T>(traverse: Traverser<T>, node?: BinaryTreeNode<T>) {
    if (!node) {
      return;
    }

    BinaryTreeNode.traverseInOrder(traverse, node.left);
    traverse(node.value);
    BinaryTreeNode.traverseInOrder(traverse, node.right);
  }

  static traversePostOrder<T>(
    traverse: Traverser<T>,
    node?: BinaryTreeNode<T>,
  ) {
    if (!node) {
      return;
    }

    BinaryTreeNode.traversePostOrder(traverse, node.left);
    BinaryTreeNode.traversePostOrder(traverse, node.right);
    traverse(node.value);
  }
}

export default BinaryTreeNode;
export { Traverser };
