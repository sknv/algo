const Queue = require('./Queue.js')

const testNewQueue = () => {
    const queue = new Queue()

    console.assert(queue.isEmpty(), 'expect queue to be empty')
}

const testQueueEnqueue = () => {
    const queue = new Queue()
    queue.enqueue(1)

    console.assert(!queue.isEmpty(), 'expect queue to have some values')
}

const testQueuePeek = () => {
    const queue = new Queue(), expected = 1
    queue.enqueue(expected)
    const value = queue.peek()

    console.assert(!queue.isEmpty(), 'expect queue to have some values')
    console.assert(value === expected, `expect to be: ${expected}, got instead: ${value}`)
}

const testQueueDequeue = () => {
    const queue = new Queue(), expected = 1
    queue.enqueue(expected)
    const value = queue.dequeue()

    console.assert(queue.isEmpty(), 'expect queue to be empty')
    console.assert(value === expected, `expect to be: ${expected}, got instead: ${value}`)
}

testNewQueue()
testQueueEnqueue()
testQueuePeek()
testQueueDequeue()
