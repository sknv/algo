mod p0217_contains_duplicate;
use p0217_contains_duplicate::Solution;

fn main() {
    let result = Solution::contains_duplicate(vec![1,2,3,1]);
    println!("result is: {:?}", result)
}
