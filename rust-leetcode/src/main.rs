mod p0049_group_anagrams;
use p0049_group_anagrams::Solution;

fn main() {
    let result = Solution::group_anagrams(vec![
        "eat".to_string(),
        "tea".to_string(),
        "tan".to_string(),
        "ate".to_string(),
        "nat".to_string(),
        "bat".to_string(),
    ]);
    println!("result is: {:?}", result)
}
