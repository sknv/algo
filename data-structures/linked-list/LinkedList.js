const LinkedListNode = require('./LinkedListNode.js')

class LinkedList {
    constructor() {
        /** @var LinkedListNode */
        this.head = null

        /** @var LinkedListNode */
        this.tail = null
    }

    /**
     * Append a new element to the tail.
     * @param {*} value
     * @returns {LinkedList}
     */
    append(value) {
        const newNode = new LinkedListNode(value)

        // If there is no head yet let's make new node a head
        if (!this.head) {
            this.head = newNode
            this.tail = newNode
            return this
        }

        // Attach new node to the end of linked list
        this.tail.next = newNode
        this.tail = newNode
        return this
    }

    /**
     * Remove and return the head.
     * @returns {LinkedListNode}
     */
    deleteHead() {
        if (!this.head) {
            return null
        }

        const deletedHead = this.head

        if (this.head.next) {
            this.head = this.head.next
        } else {
            this.head = null
            this.tail = null
        }
        return deletedHead
    }
}

module.exports = LinkedList
