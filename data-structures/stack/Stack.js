import LinkedList from '../linked-list/LinkedList.js'

class Stack {
    /** @type {LinkedList} */
    linkedList = new LinkedList()

    /**
     * @returns {boolean}
     */
    isEmpty() {
        return !this.linkedList.head
    }

    /**
     * Read the element at the front of the stack without removing it.
     * @returns {*}
     */
    peek() {
        return this.linkedList.head?.value
    }

    /**
     * Add a new element to the front of the stack (the head of the linked list).
     * This element will be processed first.
     * @param {*} value
     */
    push(value) {
        this.linkedList.prepend(value)
    }

    /**
     * Remove and return the element at the front of the queue (the head of the linked list).
     * @returns {*}
     */
    pop() {
        const deletedHead = this.linkedList.deleteHead()
        return deletedHead?.value
    }
}

export default Stack
