// https://leetcode.com/problems/top-k-frequent-elements/

use std::collections::HashMap;

pub struct Solution;

struct Pair(i32, i32); // value, frequency

impl Solution {
    // O(n * log n) - using sorting.
    // TODO: implement O(n * log k) using heap/priority queue.
    pub fn top_k_frequent(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let mut freqs = HashMap::with_capacity(nums.len());
        for num in nums {
            *freqs.entry(num).or_insert(0) += 1;
        }

        let mut pair_freqs = Vec::with_capacity(freqs.len());
        for (&val, &freq) in &freqs {
            pair_freqs.push(Pair(val, freq));
        }

        // Sort by desc.
        pair_freqs.sort_by(|prev, next| next.1.cmp(&prev.1));

        let mut result = Vec::with_capacity(k as usize);
        for i in 0..k {
            result.push(pair_freqs[i as usize].0);
        }

        result
    }
}
