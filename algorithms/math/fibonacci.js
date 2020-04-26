/**
 * @param {number} n
 * @returns {number}
 */
const fibonacci = n => {
    let cur = 0, next = 1
    for (let i = 0; i < n; i++) {
        const tmp = cur
        cur = next
        next = tmp + next
    }
    return cur
}

module.exports = fibonacci
