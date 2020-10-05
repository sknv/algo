import binarySearch from "./binarySearch.ts";

function testBinarySearch() {
  const nums = [-2, 0, 5, 9, 12];

  let idx = binarySearch(nums, 12), expected = 4;
  console.assert(
    idx === expected,
    `expect binarySearch(..., 12) to be: ${expected}, got instead: ${idx}`,
  );

  idx = binarySearch(nums, 6), expected = -1;
  console.assert(
    idx === expected,
    `expect binarySearch(..., 6) to be: ${expected}, got instead: ${idx}`,
  );
}

testBinarySearch();
