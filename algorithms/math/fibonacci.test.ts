import fibonacci from "./fibonacci.ts";

function testFibonacci() {
  let num = fibonacci(1), expected = 1;
  console.assert(
    num === expected,
    `expect fibonacci(1) to be: ${expected}, got instead: ${num}`,
  );

  num = fibonacci(2), expected = 1;
  console.assert(
    num === expected,
    `expect fibonacci(2) to be: ${expected}, got instead: ${num}`,
  );

  num = fibonacci(6), expected = 8;
  console.assert(
    num === expected,
    `expect fibonacci(6) to be: ${expected}, got instead: ${num}`,
  );
}

testFibonacci();
