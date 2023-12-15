package blockchain

import (
	"bytes"
	"crypto/sha256"
	"log"
)

type MerkleTree struct {
	RootNode *MerkleNode
}

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Hash  []byte
}

func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	node := MerkleNode{}

	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		node.Hash = hash[:]
	} else {
		prevHashes := append(left.Hash, right.Hash...)
		hash := sha256.Sum256(prevHashes)
		node.Hash = hash[:]
	}

	node.Left = left
	node.Right = right

	return &node
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode

	for _, dat := range data {
		node := NewMerkleNode(nil, nil, dat)
		nodes = append(nodes, *node)
	}

	if len(nodes) == 0 {
		log.Panic("No merkel nodes")
	}

	for len(nodes) > 1 {
		if len(nodes)%2 != 0 {
			nodes = append(nodes, nodes[len(nodes)-1])
		}

		var level []MerkleNode
		for i := 0; i < len(nodes); i += 2 {
			node := NewMerkleNode(&nodes[i], &nodes[i+1], nil)
			level = append(level, *node)
		}

		nodes = level
	}

	tree := MerkleTree{&nodes[0]}

	return &tree
}

// GetMerklePath returns the Merkle path and indices (left or right) for a given leaf node.
func (mt *MerkleTree) GetMerklePath(data []byte) (path [][]byte, indices []int) {

	hash := sha256.Sum256(data)
	var getPath func(node *MerkleNode) bool
	getPath = func(node *MerkleNode) bool {
		if node == nil {
			return false
		}
		// Leaf node
		if node.Left == nil && node.Right == nil {
			if bytes.Equal(node.Hash, hash[:]) {
				return true
			}
			return false
		}

		if getPath(node.Left) {
			path = append(path, node.Left.Hash)
			indices = append(indices, 1)
			return true
		}

		if getPath(node.Right) {
			path = append(path, node.Right.Hash)
			indices = append(indices, 0)
			return true
		}

		return false
	}

	getPath(mt.RootNode)

	return path, indices
}
