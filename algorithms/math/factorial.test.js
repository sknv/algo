import factorial from './factorial.js'

const testFactorial = () => {
    let num = factorial(0), expected = 1
    console.assert(num === expected, `expect to be: ${expected}, got instead: ${num}`)

    num = factorial(1), expected = 1
    console.assert(num === expected, `expect to be: ${expected}, got instead: ${num}`)

    num = factorial(2), expected = 2
    console.assert(num === expected, `expect to be: ${expected}, got instead: ${num}`)

    num = factorial(6), expected = 720
    console.assert(num === expected, `expect to be: ${expected}, got instead: ${num}`)
}

testFactorial()
