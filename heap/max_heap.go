package heap

import (
	"errors"
	"fmt"
)

// Elem ..
type Elem interface {
	Compare(j Elem) bool
}

// Person ..
type Person struct {
	Name  string
	Money float64
}

// Compare ..
func (s *Person) Compare(j Elem) bool {
	return s.Money > j.(*Person).Money
}

// MaxHeap ..
type MaxHeap struct {
	Heap []Elem
	Size int
}

// NewMaxHeap ..
func NewMaxHeap() *MaxHeap {
	m := &MaxHeap{}
	m.Heap = append(m.Heap, nil)
	return m
}

// Swap ..
func (m *MaxHeap) Swap(i, j int) {
	m.Heap[i], m.Heap[j] = m.Heap[j], m.Heap[i]
}

// Insert ..
func (m *MaxHeap) Insert(e Elem) {
	m.Heap = append(m.Heap, e)
	m.Size++
	m.swim(e)
}

func (m *MaxHeap) swim(e Elem) {
	k := m.Size
	// 到根节点结束
	for k/2 > 0 {
		// 判断当前节点是否大于其父节点
		if m.Heap[k].Compare(m.Heap[k/2]) {
			m.Swap(k, k/2)
		}

		k = k / 2
	}
}

// Print ..
func (m *MaxHeap) Print() {
	fmt.Printf("当前堆大小: %d\n", m.Size)
	for _, e := range m.Heap {
		fmt.Printf("%v, ", e)
	}

	fmt.Println()
}

var errEmpty = errors.New("堆为空")

// Delete 删除堆顶元素
func (m *MaxHeap) Delete() (Elem, error) {
	if len(m.Heap) == 0 {
		return nil, errEmpty
	}

	// 堆顶元素与最后一个元素交换，之后下沉调整
	m.Swap(1, m.Size)

	pop := m.Heap[m.Size]
	m.Heap[m.Size] = nil
	m.Size--

	m.sink()
	return pop, nil
}

func (m *MaxHeap) sink() {
	k := 1

	for 2*k+1 <= m.Size {
		// 初始为左节点下标
		max := 2 * k
		// 有右节点，比较并且右节点大于左节点
		if (2*k+1 <= m.Size) && (m.Heap[2*k+1].Compare(m.Heap[max])) {
			max = 2*k + 1
		}
		// max节点大于父节点，则交换，将父节点下沉
		if m.Heap[max].Compare(m.Heap[k]) {
			m.Swap(max, k)
		} else {
			break
		}

		k = max
	}
}
