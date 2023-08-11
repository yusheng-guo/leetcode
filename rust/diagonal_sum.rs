// 矩阵对角线元素的和
// https://leetcode.cn/problems/matrix-diagonal-sum/submissions/

impl Solution {
    pub fn diagonal_sum(mat: Vec<Vec<i32>>) -> i32 {
    let n = mat.len();
    let mut sum = 0;

    for i in 0..n {
        for j in 0..n {
            if i == j || i + j == n - 1 {
                sum += mat[i][j];
            }
        }
    }

    sum
    }
}