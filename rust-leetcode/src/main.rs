mod p0347_top_k_frequent_elements;
use p0347_top_k_frequent_elements::Solution;

fn main() {
    let result = Solution::top_k_frequent(vec![1,1,1,2,2,3], 2);
    println!("result is: {:?}", result)
}
