import LinkedListNode from "./LinkedListNode.ts";

class LinkedList<T> {
  head?: LinkedListNode<T>;
  tail?: LinkedListNode<T>;

  get isEmpty(): boolean {
    return !this.head;
  }

  /**
     * Return the prepended node.
     */
  prepend(value: T): LinkedListNode<T> {
    // Make new node to be the head
    const newNode = {
      value,
      next: this.head,
    };
    this.head = newNode;

    // If there is no tail yet let's make new node a tail
    if (!this.tail) {
      this.tail = newNode;
    }
    return newNode;
  }

  /**
     * Return the appended node.
     */
  append(value: T): LinkedListNode<T> {
    const newNode = { value };

    // If there is no elements yet let's make new node a head
    if (this.isEmpty) {
      this.head = newNode;
      this.tail = newNode;
      return newNode;
    }

    // Attach new node to the end of linked list
    this.tail!.next = newNode;
    this.tail = newNode;
    return newNode;
  }

  /**
     * Remove and return the head.
     */
  deleteHead(): LinkedListNode<T> | undefined {
    if (!this.head) {
      return undefined;
    }

    const deletedHead = this.head;

    if (this.head.next) {
      this.head = this.head.next;
    } else {
      this.head = undefined;
      this.tail = undefined;
    }
    return deletedHead;
  }

  /**
     * Remove and return the tail.
     */
  deleteTail(): LinkedListNode<T> | undefined {
    const deletedTail = this.tail;

    if (this.head === this.tail) { // there is no or only one node in linked list
      this.head = undefined;
      this.tail = undefined;
      return deletedTail;
    }

    // Rewind to the last node and delete "next" link for the node before the last one
    let currentNode = this.head;
    while (currentNode?.next) {
      if (!currentNode.next.next) {
        currentNode.next = undefined;
      } else {
        currentNode = currentNode.next;
      }
    }
    this.tail = currentNode;
    return deletedTail;
  }
}

export default LinkedList;
