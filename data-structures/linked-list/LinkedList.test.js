const LinkedList = require('./LinkedList.js')

const testNewLinkedList = () => {
    const linkedList = new LinkedList()

    console.assert(!linkedList.head, 'expect linked list head to be null')
    console.assert(!linkedList.tail, 'expect linked list tail to be null')
}

const testLinkedListAppend = () => {
    const linkedList = new LinkedList()
    linkedList.append(1)

    console.assert(linkedList.head, 'expect head to present')
    console.assert(linkedList.tail, 'expect tail to present')
}

const testLinkedListDeleteHead = () => {
    const linkedList = new LinkedList()
    linkedList.append(1)
    linkedList.deleteHead()

    console.assert(!linkedList.head, 'expect head to be null')
    console.assert(!linkedList.tail, 'expect tail to be null')
}

testNewLinkedList()
testLinkedListAppend()
testLinkedListDeleteHead()
