// https://leetcode.com/problems/number-of-islands/

pub struct Solution;

impl Solution {
    pub fn num_islands(mut grid: Vec<Vec<char>>) -> i32 {
        let mut islands = 0;

        let rows = grid.len() as i32;
        let cols = grid[0].len() as i32;

        for i in 0..rows {
            for j in 0..cols {
                if grid[i as usize][j as usize] == '1' {
                    dfs(&mut grid, i, j);
                    islands += 1;
                }
            }
        }

        islands
    }
}

fn dfs(grid: &mut Vec<Vec<char>>, row: i32, col: i32) {
    let rows = grid.len() as i32;
    let cols = grid[0].len() as i32;

    if row < 0 || col < 0 || row >= rows || col >= cols || grid[row as usize][col as usize] == '0' {
        return;
    }

    grid[row as usize][col as usize] = '0'; // Visit the point.
    dfs(grid, row - 1, col);
    dfs(grid, row + 1, col);
    dfs(grid, row, col - 1);
    dfs(grid, row, col + 1);
}
