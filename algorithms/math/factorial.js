/**
 * @param {number} n
 * @returns {number}
 */
const factorial = n => {
    let result = 1
    for (let i = 2; i <= n; i++) {
        result *= i
    }
    return result
}

export default factorial
