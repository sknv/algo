mod p0695_max_area_of_island;
use p0695_max_area_of_island::Solution;

fn main() {
    let result = Solution::max_area_of_island(vec![vec![1, 1, 1, 1, 0]]);
    println!("result is: {:?}", result)
}
