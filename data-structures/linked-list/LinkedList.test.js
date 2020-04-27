const LinkedList = require('./LinkedList.js')

const testNewLinkedList = () => {
    const list = new LinkedList()

    console.assert(list.head === null, 'expect list head to be null')
    console.assert(list.tail === null, 'expect list tail to be null')
}

const testLinkedListAppend = () => {
    const list = new LinkedList()
    list.append(1)

    console.assert(list.head !== null, 'expect head to present')
    console.assert(list.tail !== null, 'expect tail to present')
}

const testLinkedDeleteHead = () => {
    const list = new LinkedList()
    list.append(1)
    list.deleteHead()

    console.assert(list.head === null, 'expect head to be null')
    console.assert(list.tail === null, 'expect tail to be null')
}

testNewLinkedList()
testLinkedListAppend()
testLinkedDeleteHead()
