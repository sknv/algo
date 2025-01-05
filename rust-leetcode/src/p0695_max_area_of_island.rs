// https://leetcode.com/problems/max-area-of-island/

pub struct Solution;

impl Solution {
    pub fn max_area_of_island(mut grid: Vec<Vec<i32>>) -> i32 {
        let mut max_area = 0;

        let rows = grid.len() as i32;
        let cols = grid[0].len() as i32;

        for i in 0..rows {
            for j in 0..cols {
                if grid[i as usize][j as usize] == 1 {
                    let cur_area = dfs(&mut grid, i, j);
                    max_area = max_area.max(cur_area);
                }
            }
        }

        max_area
    }
}

fn dfs(grid: &mut Vec<Vec<i32>>, row: i32, col: i32) -> i32 {
    let rows = grid.len() as i32;
    let cols = grid[0].len() as i32;

    if row < 0 || col < 0 || row >= rows || col >= cols || grid[row as usize][col as usize] == 0 {
        return 0;
    }

    grid[row as usize][col as usize] = 0; // Visit the point.
    1 + dfs(grid, row - 1, col)
        + dfs(grid, row + 1, col)
        + dfs(grid, row, col - 1)
        + dfs(grid, row, col + 1)
}
