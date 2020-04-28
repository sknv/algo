/**
 * @param {number} n
 * @returns {number}
 */
const fibonacci = n => {
    let cur = 0, next = 1
    for (let i = 0; i < n; i++) {
        [cur, next] = [next, cur + next] // fancy ES6 swap
    }
    return cur
}

export default fibonacci
