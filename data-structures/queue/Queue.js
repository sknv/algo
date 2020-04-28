const LinkedList = require('../linked-list/LinkedList.js')

class Queue {
    /** @type {LinkedList} */
    linkedList = new LinkedList()

    /**
     * @returns {boolean}
     */
    isEmpty() {
        return !this.linkedList.head
    }

    /**
     * Read the element at the front of the queue without removing it.
     * @returns {*}
     */
    peek() {
        return this.linkedList.head?.value
    }

    /**
     * Add a new element to the end of the queue (the tail of the linked list).
     * This element will be processed after all elements ahead of it.
     * @param {*} value
     */
    enqueue(value) {
        this.linkedList.append(value)
    }

    /**
     * Remove the element at the front of the queue (the head of the linked list).
     * @returns {*}
     */
    dequeue() {
        const deletedHead = this.linkedList.deleteHead()
        return deletedHead?.value
    }
}

module.exports = Queue
