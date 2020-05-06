function factorial(n: number): number {
    let result = 1
    for (let i = 2; i <= n; i++) {
        result *= i
    }
    return result
}

export default factorial
