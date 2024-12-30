// https://leetcode.com/problems/valid-anagram/

use std::collections::HashMap;

pub struct Solution;

impl Solution {
    pub fn is_anagram(s: String, t: String) -> bool {
        if s.len() != t.len() {
            return false;
        }

        let mut s_chars_count = HashMap::with_capacity(s.len());
        for chr in s.chars() {
            *s_chars_count.entry(chr).or_insert(0) += 1;
        }

        for chr in t.chars() {
            let s_char = s_chars_count.entry(chr).or_insert(0);
            *s_char -= 1;

            if *s_char < 0 {
                return false;
            }
        }

        true
    }
}
