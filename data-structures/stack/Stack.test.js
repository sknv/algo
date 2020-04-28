import Stack from './Stack.js'

const testNewStack = () => {
    const stack = new Stack()

    console.assert(stack.isEmpty(), 'expect stack to be empty')
}

const testStackPush = () => {
    const stack = new Stack()
    stack.push(1)

    console.assert(!stack.isEmpty(), 'expect stack to have some values')
}

const testStackPeek = () => {
    const stack = new Stack(), expected = 1
    stack.push(expected)
    const value = stack.peek()

    console.assert(!stack.isEmpty(), 'expect stack to have some values')
    console.assert(value === expected, `expect to be: ${expected}, got instead: ${value}`)
}

const testStackPop = () => {
    const stack = new Stack(), expected = 2
    stack.push(1)
    stack.push(expected)

    const value = stack.pop()
    console.assert(value === expected, `expect to be: ${expected}, got instead: ${value}`)

    stack.pop()
    console.assert(stack.isEmpty(), 'expect stack to be empty')
}

testNewStack()
testStackPush()
testStackPeek()
testStackPop()
