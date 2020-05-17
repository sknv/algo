import BinaryTreeNode, { Traverser } from "./BinaryTreeNode.ts";

class BinaryTree<T> {
  constructor(public root?: BinaryTreeNode<T>) {}

  get leftHeight(): number {
    return BinaryTreeNode.getLeftHeight(this.root);
  }

  get rightHeight(): number {
    return BinaryTreeNode.getRightHeight(this.root);
  }

  get height(): number {
    return BinaryTreeNode.getHeight(this.root);
  }

  get balanceFactor(): number {
    return BinaryTreeNode.getBalanceFactor(this.root);
  }

  traversePreOrder(traverse: Traverser<T>) {
    BinaryTreeNode.traversePreOrder(traverse, this.root);
  }

  traverseInOrder(traverse: Traverser<T>) {
    BinaryTreeNode.traverseInOrder(traverse, this.root);
  }

  traversePostOrder(traverse: Traverser<T>) {
    BinaryTreeNode.traversePostOrder(traverse, this.root);
  }
}

export default BinaryTree;
