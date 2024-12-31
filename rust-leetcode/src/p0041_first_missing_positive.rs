// https://leetcode.com/problems/first-missing-positive/

pub struct Solution;

impl Solution {
    pub fn first_missing_positive(mut nums: Vec<i32>) -> i32 {
        let count = nums.len();

        // Step 1: Mark present positive integers.
        for i in 0..count {
            let mut num = nums[i];
            while num > 0 && num <= count as i32 && nums[(num - 1) as usize] != num {
                nums.swap(i, (num - 1) as usize); // Swap nums[i] with nums[num - 1] to place num in its correct position.
                num = nums[i];
            }
        }

        // Step 2: Find the smallest missing positive integer.
        for i in 0..count {
            if nums[i] != (i + 1) as i32 {
                return (i + 1) as i32;
            }
        }

        // If all numbers from 1 to count are present, return count + 1.
        (count + 1) as i32
    }
}
