mod p0268_missing_number;
use p0268_missing_number::Solution;

fn main() {
    let result = Solution::missing_number(vec![3, 0, 1]);
    println!("result is: {:?}", result)
}
