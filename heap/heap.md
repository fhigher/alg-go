## 堆 heap

### 特性
    1. 堆中任一节点的值总是小于等于其父节点的值
    2. 堆是一棵完全二叉数
    3. 最大堆，最小堆
    4. 堆可以用于实现优先队列

### 实现优先队列优势比较
|   | 入队  | 出队  |
|---|---|---|
| 普通数组  | O(1) | O(n) |
| 顺序数组  | O(n) | O(1) |
| 堆  | O(logn) | O(logn) |

### 最差性能
    普通数组和顺序数组会演变为O(n^2)
    堆则演变为O(nlogn)

### 数组存储二叉堆
    从下标1开始存储，则
    parent(i) = i / 2
    left child(i) = 2 * i
    right child(i) = 2 * i + 1