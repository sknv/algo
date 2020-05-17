import LinkedList from "../linked-list/LinkedList.ts";

class Stack<T> {
  linkedList = new LinkedList<T>();

  get isEmpty(): boolean {
    return this.linkedList.isEmpty;
  }

  /**
     * Read the element at the front of the stack without removing it.
     */
  peek(): T | undefined {
    return this.linkedList.head?.value;
  }

  /**
     * Add a new element to the front of the stack (the head of the linked list).
     * This element will be processed first.
     */
  push(value: T) {
    this.linkedList.prepend(value);
  }

  /**
     * Remove and return the element at the front of the queue (the head of the linked list).
     */
  pop(): T | undefined {
    const deletedHead = this.linkedList.deleteHead();
    return deletedHead?.value;
  }
}

export default Stack;
