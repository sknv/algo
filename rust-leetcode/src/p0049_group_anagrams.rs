// https://leetcode.com/problems/group-anagrams/

use std::collections::HashMap;

pub struct Solution;

impl Solution {
    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
        let mut group_freqs = HashMap::with_capacity(strs.len());
        for str in strs {
            let char_freqs = count_char_frequencies(&str);
            group_freqs.entry(char_freqs).or_insert(vec![]).push(str);
        }

        group_freqs.into_values().collect()
    }
}

fn count_char_frequencies(str: &str) -> [u8; 26] {
    let mut freqs: [u8; 26] = [0; 26];
    for chr in str.chars() {
        let idx = chr as u8 - 'a' as u8;
        freqs[idx as usize] += 1;
    }

    freqs
}
