package heap

import (
	"errors"
	"fmt"
)

// Elem ..
type Elem interface {
	Compare(j Elem) bool
}

// MaxHeap ..
type MaxHeap struct {
	heap []Elem
	size int
}

// NewMaxHeap ..
func NewMaxHeap() *MaxHeap {
	m := &MaxHeap{}
	m.heap = append(m.heap, nil)
	return m
}

// Swap ..
func (m *MaxHeap) Swap(i, j int) {
	m.heap[i], m.heap[j] = m.heap[j], m.heap[i]
}

// Insert ..
func (m *MaxHeap) Insert(e Elem) {
	m.heap = append(m.heap, e)
	m.size++
	m.swim(m.size)
}

func (m *MaxHeap) FromArray(arr []Elem) {
	for i := 0; i < len(arr); i++ {
		m.heap = append(m.heap, arr[i])
	}
	m.size = len(arr)

	for i := len(arr) / 2; i >= 1; i-- {
		m.sink(i)
	}
}

func (m *MaxHeap) Size() int {
	return m.size
}

func (m *MaxHeap) swim(k int) {
	// 到根节点结束
	for k/2 > 0 {
		// 判断当前节点是否大于其父节点
		if m.heap[k].Compare(m.heap[k/2]) {
			m.Swap(k, k/2)
		}

		k = k / 2
	}
}

// Print ..
func (m *MaxHeap) Print() {
	fmt.Printf("当前堆大小: %d\n", m.size)
	for _, e := range m.heap {
		fmt.Printf("%v, ", e)
	}

	fmt.Println()
}

var errEmpty = errors.New("堆为空")

// Delete 删除堆顶元素
func (m *MaxHeap) DeleteMax() (Elem, error) {
	if len(m.heap) == 0 {
		return nil, errEmpty
	}

	// 堆顶元素与最后一个元素交换，之后下沉调整
	m.Swap(1, m.size)

	pop := m.heap[m.size]
	m.heap[m.size] = nil
	m.size--

	m.sink(1)
	return pop, nil
}

func (m *MaxHeap) sink(k int) {

	for 2*k+1 <= m.size {
		// 初始为左节点下标
		max := 2 * k
		// 有右节点，比较并且右节点大于左节点
		if (2*k+1 <= m.size) && (m.heap[2*k+1].Compare(m.heap[max])) {
			max = 2*k + 1
		}
		// max节点大于父节点，则交换，将父节点下沉
		if m.heap[max].Compare(m.heap[k]) {
			m.Swap(max, k)
		} else {
			break
		}

		k = max
	}
}
