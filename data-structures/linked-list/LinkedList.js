class LinkedList {
    /** @type {LinkedListNode} */
    head

    /** @type {LinkedListNode} */
    tail

    /**
     * @param {*} value
     * @returns {LinkedList}
     */
    prepend(value) {
        // Make new node to be the head
        /** @type {LinkedListNode} */
        const newNode = { value, next: this.head }
        this.head = newNode

        // If there is no tail yet let's make new node a tail
        if (!this.tail) {
            this.tail = newNode
        }
        return this
    }

    /**
     * Append a new element to the tail.
     * @param {*} value
     * @returns {LinkedList}
     */
    append(value) {
        /** @type {LinkedListNode} */
        const newNode = { value }

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

export default LinkedList
