package tree

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
)

// 交易结构
type Transcation = []byte

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

func newMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {

	var (
		node     = &MerkleNode{}
		hash     [32]byte
		byteHash []byte
	)

	if left == nil && right == nil {
		// 叶子节点存储元数据
		byteHash = data
	} else {
		// 父节点存储元数据
		byteHash = append(left.Data, right.Data...)
	}
	// 节点存储hash值
	hash = sha256.Sum256(byteHash)
	node.Data = hash[:]

	node.Left = left
	node.Right = right

	return node
}

type MerkleTree struct {
	root *MerkleNode
}

func NewMerkleTree(txs ...Transcation) *MerkleTree {
	// 检查交易数据是否为偶数，不是则复制最后一笔交易
	if len(txs)%2 != 0 {
		txs = append(txs, txs[len(txs)-1])
	}

	// 遍历交易组，第一层使用交易数据一一对应创建叶子节点。
	var nodes []*MerkleNode
	for i := 0; i < len(txs); i++ {
		nodes = append(nodes, newMerkleNode(nil, nil, txs[i]))
	}
	fmt.Println(len(nodes))
	// 之后每一层，节点数量减半
	// 第二层及之后层(不包含最顶层)，节点数不足偶数时，复制最后一个节点.Data，左右子节点创建其父节点
	// 直到最顶层生成根节点
	// i 为树的层编号, j 为每一层节点的编号
	var pNode *MerkleNode
	for i := 0; i < len(txs)/2; i++ {
		var parentNodes []*MerkleNode
		for j := 0; j < len(nodes); j += 2 {
			// 被复制的节点也会添加进树里面
			pNode = newMerkleNode(nodes[j], nodes[j+1], nil)
			parentNodes = append(parentNodes, pNode)
		}

		parentLength := len(parentNodes)
		if parentLength%2 != 0 && parentLength != 1 {
			parentNodes = append(parentNodes, parentNodes[parentLength-1])
		}

		nodes = parentNodes
		fmt.Println(len(nodes))
		if len(nodes) == 1 {
			break
		}
	}

	return &MerkleTree{nodes[0]}
}

// 前序遍历
func (mt *MerkleTree) PreOrderRange(root *MerkleNode) {
	if root == nil {
		return
	}

	fmt.Printf("节点hash为：%s, ", hex.EncodeToString(root.Data))

	if root.Left == nil || root.Right == nil {
		log.Println("该节点为叶子节点")
		return
	}
	hash := append(root.Left.Data, root.Right.Data...)
	hashByte := sha256.Sum256(hash)
	if bytes.Equal(root.Data, hashByte[:]) {
		log.Println("该节点的hash值, 等于两个子节点组合后的hash值")
	} else {
		log.Println("该节点的hash值, 不等于两个子节点组合后的hash值")
		return
	}

	mt.PreOrderRange(root.Left)
	// 打印右侧的分支时，如果那一层的节点当时为奇数，从而被复制了一次。那么被复制的节点及其子节点会打印两次
	mt.PreOrderRange(root.Right)
}
