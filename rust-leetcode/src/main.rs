mod p0200_number_of_islands;
use std::vec;

use p0200_number_of_islands::Solution;

fn main() {
    let result = Solution::num_islands(vec![vec!['1', '1', '1', '1', '0']]);
    println!("result is: {:?}", result)
}
