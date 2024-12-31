mod p0041_first_missing_positive;
use p0041_first_missing_positive::Solution;

fn main() {
    let result = Solution::first_missing_positive(vec![3, 4, -1, 1]);
    println!("result is: {:?}", result)
}
