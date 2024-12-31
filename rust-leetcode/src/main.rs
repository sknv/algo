mod p0448_array_numbers_disappeared;
use p0448_array_numbers_disappeared::Solution;

fn main() {
    let result = Solution::find_disappeared_numbers(vec![4,3,2,7,8,2,3,1]);
    println!("result is: {:?}", result)
}
