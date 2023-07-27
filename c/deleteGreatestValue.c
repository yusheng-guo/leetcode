int cmp(void const* a, void const* b) {
    return *(int*)a - *(int*)b;
}

int deleteGreatestValue(int** grid, int gridSize, int* gridColSize) {
    int m = gridSize, n = *gridColSize;
    for (int i = 0; i < m; ++i) {
        qsort(grid[i], n, 4U, cmp);
    }
    int res = 0;
    for (int j = 0; j < n; j++) {
        int mx = 0;
        for (int i = 0; i < m; ++i) {
            if (grid[i][j] > mx) {
                mx = grid[i][j];
            }
        }
        res += mx;
    }
    return res;
}