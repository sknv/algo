import bubbleSort from "./bubbleSort.ts";

function testBubbleSort() {
  let nums = [5, -2, 12, 0, 9], expected = [-2, 0, 5, 9, 12];
  bubbleSort(nums);
  console.assert(
    JSON.stringify(nums) === JSON.stringify(expected),
    `expect bubbleSort(...) to be: ${expected}, got instead: ${nums}`,
  );

  nums = [1, -2, 10, -6], expected = [-6, -2, 1, 10];
  bubbleSort(nums);
  console.assert(
    JSON.stringify(nums) === JSON.stringify(expected),
    `expect bubbleSort(...) to be: ${expected}, got instead: ${nums}`,
  );
}

testBubbleSort();
