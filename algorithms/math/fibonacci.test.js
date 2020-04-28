import fibonacci from './fibonacci.js'

const testFibonacci = () => {
    let num = fibonacci(1), expected = 1
    console.assert(num === expected, `expect to be: ${expected}, got instead: ${num}`)

    num = fibonacci(2), expected = 1
    console.assert(num === expected, `expect to be: ${expected}, got instead: ${num}`)

    num = fibonacci(6), expected = 8
    console.assert(num === expected, `expect to be: ${expected}, got instead: ${num}`)
}

testFibonacci()
