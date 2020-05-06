import Stack from './Stack.ts'

function testNewStack() {
    const stack = new Stack<number>()

    console.assert(stack.isEmpty, 'expect Stack<> to be empty')
}

function testStackPush() {
    const stack = new Stack<number>()
    stack.push(1)

    console.assert(!stack.isEmpty, 'expect Stack<> to have some values')
}

function testStackPeek() {
    const stack = new Stack<number>(), expected = 1
    stack.push(expected)
    const value = stack.peek()

    console.assert(!stack.isEmpty, 'expect Stack<> to have some values')
    console.assert(value === expected, `expect to be: ${expected}, got instead: ${value}`)
}

function testStackPop() {
    const stack = new Stack<number>(), expected = 2
    stack.push(1)
    stack.push(expected)

    const value = stack.pop()
    console.assert(value === expected, `expect to be: ${expected}, got instead: ${value}`)

    stack.pop()
    console.assert(stack.isEmpty, 'expect Stack<> to be empty')
}

testNewStack()
testStackPush()
testStackPeek()
testStackPop()
