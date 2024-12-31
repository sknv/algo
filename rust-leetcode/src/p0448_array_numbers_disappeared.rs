// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array

pub struct Solution;

impl Solution {
    pub fn find_disappeared_numbers(mut nums: Vec<i32>) -> Vec<i32> {
        let count = nums.len();
        for i in 0..count {
            let idx = (nums[i].abs() - 1) as usize;
            if nums[idx] > 0 {
                nums[idx] = -nums[idx]
            }
        }

        nums.into_iter()
            .enumerate()
            .filter(|&(_, val)| val > 0)
            .map(|(idx, _)| (idx+1) as i32)
            .collect()
    }
}
