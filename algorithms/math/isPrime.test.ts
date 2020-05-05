import isPrime from './isPrime.ts'

function testIsPrime() {
    let prime = isPrime(1), expected = false
    console.assert(prime === expected, `expect isPrime(1) to be: ${expected}, got instead: ${prime}`)

    prime = isPrime(2), expected = true
    console.assert(prime === expected, `expect isPrime(2) to be: ${expected}, got instead: ${prime}`)

    prime = isPrime(3), expected = true
    console.assert(prime === expected, `expect isPrime(3) to be: ${expected}, got instead: ${prime}`)

    prime = isPrime(4), expected = false
    console.assert(prime === expected, `expect isPrime(4) to be: ${expected}, got instead: ${prime}`)

    prime = isPrime(6), expected = false
    console.assert(prime === expected, `expect isPrime(6) to be: ${expected}, got instead: ${prime}`)

    prime = isPrime(11), expected = true
    console.assert(prime === expected, `expect isPrime(11) to be: ${expected}, got instead: ${prime}`)

    prime = isPrime(97), expected = true
    console.assert(prime === expected, `expect isPrime(97) to be: ${expected}, got instead: ${prime}`)
}

testIsPrime()
