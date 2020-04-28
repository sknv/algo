import isPrime from './isPrime.js'

const testIsPrime = () => {
    let prime = isPrime(1), expected = false
    console.assert(prime === expected, `expect to be: ${expected}, got instead: ${prime}`)

    prime = isPrime(2), expected = true
    console.assert(prime === expected, `expect to be: ${expected}, got instead: ${prime}`)

    prime = isPrime(3), expected = true
    console.assert(prime === expected, `expect to be: ${expected}, got instead: ${prime}`)

    prime = isPrime(4), expected = false
    console.assert(prime === expected, `expect to be: ${expected}, got instead: ${prime}`)

    prime = isPrime(6), expected = false
    console.assert(prime === expected, `expect to be: ${expected}, got instead: ${prime}`)

    prime = isPrime(11), expected = true
    console.assert(prime === expected, `expect to be: ${expected}, got instead: ${prime}`)

    prime = isPrime(97), expected = true
    console.assert(prime === expected, `expect to be: ${expected}, got instead: ${prime}`)
}

testIsPrime()
