// https://leetcode.com/problems/coin-change/description/

use std::collections::HashMap;

pub struct Solution;

impl Solution {
    pub fn coin_change(coins: Vec<i32>, amount: i32) -> i32 {
        let mut memo = HashMap::with_capacity(coins.len()); // Memoization table
        let result = dp(&coins, amount, &mut memo);

        if result == i32::MAX {
            -1 // If the result is still i32::MAX, the amount cannot be made up
        } else {
            result
        }
    }
}

fn dp(coins: &Vec<i32>, amount: i32, memo: &mut HashMap<i32, i32>) -> i32 {
    if amount == 0 {
        return 0; // Base case: 0 coins are needed for amount 0
    }

    if amount < 0 {
        return i32::MAX; // Invalid case: return a large value to indicate impossibility
    }

    // Check if the result is already computed
    if let Some(&result) = memo.get(&amount) {
        return result;
    }

    let mut min_coins = i32::MAX;

    // Try each coin
    for &coin in coins {
        let res = dp(coins, amount - coin, memo);
        if res != i32::MAX {
            min_coins = min_coins.min(res + 1);
        }
    }

    // Store the result in the memoization table
    memo.insert(amount, min_coins);

    min_coins
}
