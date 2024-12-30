// https://leetcode.com/problems/two-sum/

use std::collections::HashMap;

pub struct Solution;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut seen = HashMap::with_capacity(nums.len());

        for i in 0..nums.len() {
            let cur_num = nums[i];
            let target_diff = target - cur_num;

            if let Some(&target_index) = seen.get(&target_diff) {
                return vec![target_index as i32, i as i32];
            }

            seen.insert(cur_num, i);
        }

        vec![]
    }
}
