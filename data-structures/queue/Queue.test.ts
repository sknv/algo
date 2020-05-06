import Queue from './Queue.ts'

function testNewQueue() {
    const queue = new Queue<number>()

    console.assert(queue.isEmpty, 'expect Queue<> to be empty')
}

function testQueueEnqueue() {
    const queue = new Queue<number>()
    queue.enqueue(1)

    console.assert(!queue.isEmpty, 'expect Queue<> to have some values')
}

function testQueuePeek() {
    const queue = new Queue<number>(), expected = 1
    queue.enqueue(expected)
    const value = queue.peek()

    console.assert(!queue.isEmpty, 'expect Queue<> to have some values')
    console.assert(value === expected, `expect Queue<>.peek() to be: ${expected}, got instead: ${value}`)
}

function testQueueDequeue() {
    const queue = new Queue<number>(), expected = 1
    queue.enqueue(expected)
    queue.enqueue(2)

    const value = queue.dequeue()
    console.assert(value === expected, `expect Queue<>.dequeue() to be: ${expected}, got instead: ${value}`)

    queue.dequeue()
    console.assert(queue.isEmpty, 'expect Queue<> to be empty')
}

testNewQueue()
testQueueEnqueue()
testQueuePeek()
testQueueDequeue()
