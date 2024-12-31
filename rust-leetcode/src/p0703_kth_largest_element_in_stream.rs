// https://leetcode.com/problems/kth-largest-element-in-a-stream/

use std::{cmp::Reverse, collections::BinaryHeap};

pub struct KthLargest {
    k: usize,
    min_heap: BinaryHeap<Reverse<i32>>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl KthLargest {
    pub fn new(k: i32, nums: Vec<i32>) -> Self {
        let k = k as usize;
        let mut min_heap = BinaryHeap::with_capacity(nums.len());

        for num in nums {
            min_heap.push(Reverse(num));
            if min_heap.len() > k {
                min_heap.pop(); // Keep the heap size <= k
            }
        }

        Self { k, min_heap }
    }

    pub fn add(&mut self, val: i32) -> i32 {
        self.min_heap.push(Reverse(val));
        if self.min_heap.len() > self.k {
            self.min_heap.pop(); // Keep the heap size <= k
        }

        if let Some(&Reverse(val)) = self.min_heap.peek() {
            val
        } else {
            0
        }
    }
}
