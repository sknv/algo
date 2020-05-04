class LinkedList {
    /** @type {LinkedListNode} */
    head

    /** @type {LinkedListNode} */
    tail

    /**
     * @param {*} value
     * @returns {LinkedListNode}
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
        return newNode
    }

    /**
     * Append a new element to the tail.
     * @param {*} value
     * @returns {LinkedListNode}
     */
    append(value) {
        /** @type {LinkedListNode} */
        const newNode = { value }

        // If there is no head yet let's make new node a head
        if (!this.head) {
            this.head = newNode
            this.tail = newNode
            return newNode
        }

        // Attach new node to the end of linked list
        this.tail.next = newNode
        this.tail = newNode
        return newNode
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

    /**
     * @returns {LinkedListNode}
     */
    deleteTail() {
        const deletedTail = this.tail

        if (this.head === this.tail) { // there is only one node in linked list
            this.head = null
            this.tail = null
            return deletedTail
        }

        // Rewind to the last node and delete "next" link for the node before the last one
        let currentNode = this.head
        while (currentNode.next) {
            if (!currentNode.next.next) {
                currentNode.next = null
            } else {
                currentNode = currentNode.next
            }
        }
        this.tail = currentNode
        return deletedTail
    }
}

export default LinkedList
