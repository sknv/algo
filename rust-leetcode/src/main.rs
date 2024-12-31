mod p0703_kth_largest_element_in_stream;
use p0703_kth_largest_element_in_stream::KthLargest;

fn main() {
    let mut obj = KthLargest::new(3, vec![4, 5, 8, 2]);
    let result = obj.add(3);
    println!("result is: {:?}", result)

    // let result = Solution::first_missing_positive(vec![3, 4, -1, 1]);
    // println!("result is: {:?}", result)
}
