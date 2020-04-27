class LinkedListNode {
    /**
     * @param {*} value
     * @param {LinkedListNode} next
     */
    constructor(value, next = null) {
        this.value = value
        this.next = next
    }

    toString() {
        return `${this.value}`
    }
}

module.exports = LinkedListNode
