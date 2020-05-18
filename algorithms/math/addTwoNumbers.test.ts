import addTwoNumbers from "./addTwoNumbers.ts";

function testAddTwoNumbers() {
  let sum = addTwoNumbers([1, 3], [4, 0, 4]), expected = [4, 1, 7];
  console.assert(
    JSON.stringify(sum) === JSON.stringify(expected),
    `expect addTwoNumbers(13, 404) to be: ${expected}, got instead: ${sum}`,
  );

  sum = addTwoNumbers([2, 1, 9], [6, 0, 1, 3, 4, 0, 4]),
    expected = [6, 0, 1, 3, 6, 2, 3];
  console.assert(
    JSON.stringify(sum) === JSON.stringify(expected),
    `expect addTwoNumbers(219, 6013404) to be: ${expected}, got instead: ${sum}`,
  );

  sum = addTwoNumbers([0], [5, 7, 8]), expected = [5, 7, 8];
  console.assert(
    JSON.stringify(sum) === JSON.stringify(expected),
    `expect addTwoNumbers(0, 578) to be: ${expected}, got instead: ${sum}`,
  );

  sum = addTwoNumbers([0], [0]), expected = [0];
  console.assert(
    JSON.stringify(sum) === JSON.stringify(expected),
    `expect addTwoNumbers(0, 0) to be: ${expected}, got instead: ${sum}`,
  );

  sum = addTwoNumbers(
    [5],
    [9, 2, 2, 3, 3, 7, 2, 0, 3, 6, 8, 5, 4, 7, 7, 5, 8, 0, 7],
  );
  expected = [9, 2, 2, 3, 3, 7, 2, 0, 3, 6, 8, 5, 4, 7, 7, 5, 8, 1, 2];
  console.assert(
    JSON.stringify(sum) === JSON.stringify(expected),
    `expect addTwoNumbers(0, 9223372036854775807) to be: ${expected}, got instead: ${sum}`,
  );
}

testAddTwoNumbers();
