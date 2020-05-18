function addTwoNumbers(a: number[], b: number[]): number[] {
  let carry = 0;
  let result: number[] = [];
  for (let i = a.length - 1, j = b.length - 1; i >= 0 || j >= 0; i--, j--) {
    let x = 0, y = 0;
    if (i >= 0) {
      x = a[i];
    }
    if (j >= 0) {
      y = b[j];
    }
    const sum = x + y + carry;
    const val = sum % 10;
    carry = Math.trunc(sum / 10);
    result.push(val);
  }
  if (carry > 0) {
    result.push(carry);
  }
  return result.reverse();
}

export default addTwoNumbers;
