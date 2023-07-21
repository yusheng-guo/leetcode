int findMaxValueOfEquation(int** points, int pointsSize, int* pointsColSize, int k) {
    int res = INT_MIN;
    int deque[pointsSize][2];
    int head = 0, tail = 0;
    for (int i = 0; i < pointsSize; i++) {
        int x = points[i][0], y = points[i][1];
        while (head != tail && x - deque[head][1] > k) {
            head++;
        }
        if (head != tail) {
            res = fmax(res, x + y + deque[head][0]);
        }
        while (head != tail && y - x >= deque[tail - 1][0]) {
            tail--;
        }
        deque[tail][0] = y - x;
        deque[tail][1] = x;
        tail++;
    }
    return res;
}