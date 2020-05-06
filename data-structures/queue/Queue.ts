import LinkedList from '../linked-list/LinkedList.ts'

class Queue<T> {
    linkedList = new LinkedList<T>()

    get isEmpty(): boolean {
        return this.linkedList.isEmpty
    }

    /**
     * Read the element at the front of the queue without removing it.
     */
    peek(): T | undefined {
        return this.linkedList.head?.value
    }

    /**
     * Add a new element to the end of the queue (the tail of the linked list).
     * This element will be processed after all elements ahead of it.
     */
    enqueue(value: T) {
        this.linkedList.append(value)
    }

    /**
     * Remove and return the element at the front of the queue (the head of the linked list).
     */
    dequeue(): T | undefined {
        const deletedHead = this.linkedList.deleteHead()
        return deletedHead?.value
    }
}

export default Queue
