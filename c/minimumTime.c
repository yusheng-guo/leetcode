typedef struct {
    int key;
    int val;
    UT_hash_handle hh;
} HashItem; 

HashItem *hashFindItem(HashItem **obj, int key) {
    HashItem *pEntry = NULL;
    HASH_FIND_INT(*obj, &key, pEntry);
    return pEntry;
}

bool hashAddItem(HashItem **obj, int key, int val) {
    if (hashFindItem(obj, key)) {
        return false;
    }
    HashItem *pEntry = (HashItem *)malloc(sizeof(HashItem));
    pEntry->key = key;
    pEntry->val = val;
    HASH_ADD_INT(*obj, key, pEntry);
    return true;
}

bool hashSetItem(HashItem **obj, int key, int val) {
    HashItem *pEntry = hashFindItem(obj, key);
    if (!pEntry) {
        hashAddItem(obj, key, val);
    } else {
        pEntry->val = val;
    }
    return true;
}

int hashGetItem(HashItem **obj, int key, int defaultVal) {
    HashItem *pEntry = hashFindItem(obj, key);
    if (!pEntry) {
        return defaultVal;
    }
    return pEntry->val;
}

void hashFree(HashItem **obj) {
    HashItem *curr = NULL, *tmp = NULL;
    HASH_ITER(hh, *obj, curr, tmp) {
        HASH_DEL(*obj, curr);  
        free(curr);
    }
}

struct ListNode *creatListNode(int val) {
    struct ListNode *obj = (struct ListNode *)malloc(sizeof(struct ListNode));
    obj->val = val;
    obj->next = NULL;
    return obj;
}

int dp(int i, const int *time, struct ListNode **prev, HashItem **memo) {
    if (!hashFindItem(memo, i)) {
        int cur = 0;
        for (struct ListNode *node = prev[i]; node != NULL; node = node->next) {
            int p = node->val;
            cur = fmax(cur, dp(p, time, prev, memo));
        }
        cur += time[i - 1];
        hashAddItem(memo, i, cur);
    }
    return hashGetItem(memo, i, 0);
}

int minimumTime(int n, int** relations, int relationsSize, int* relationsColSize, int* time, int timeSize) {
    int mx = 0;
    struct ListNode *prev[n + 1];
    for (int i = 0; i <= n; i++) {
        prev[i] = NULL;
    }
    for (int i = 0; i < relationsSize; i++) {
        int x = relations[i][0], y = relations[i][1];
        struct ListNode *node = creatListNode(x);
        node->next = prev[y];
        prev[y] = node;
    }
    HashItem *memo = NULL;
    for (int i = 1; i <= n; i++) {
        mx = fmax(mx, dp(i, time, prev, &memo));
    }
    hashFree(&memo);
    for (int i = 0; i <= n; i++) {
        struct ListNode *node = prev[i];
        while (node) {
            struct ListNode *cur = node;
            node = node->next;
            free(cur);
        }
    }
    return mx;
}
