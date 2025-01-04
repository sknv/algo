mod p0322_coin_change;
use p0322_coin_change::Solution;

fn main() {
    let result = Solution::coin_change(vec![1, 2, 5], 11);
    println!("result is: {:?}", result)
}
