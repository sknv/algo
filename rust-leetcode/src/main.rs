mod p0242_valid_anagram;
use p0242_valid_anagram::Solution;

fn main() {
    let result = Solution::is_anagram("anagram".to_string(), "nagaram".to_string());
    println!("result is: {:?}", result)
}
