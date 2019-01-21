package algorithms

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type MerkleNode struct {
	Data   []byte
	hash   []byte
	Parent *MerkleNode
	left   *MerkleNode
	right  *MerkleNode
}
type MerkleTree struct {
	leafs    []*MerkleNode
	Root     *MerkleNode
	RootHash []byte
}

//将待加密数据转换成叶子节点切片
func convert2Node(contents []string) []*MerkleNode {
	var leafs []*MerkleNode
	for i := 0; i < len(contents); i++ {
		node := MerkleNode{
			Data: []byte(contents[i]),
			hash: getHash([]byte(contents[i])),
		}
		leafs = append(leafs, &node)
	}
	return leafs
}

//初始化梅克尔根
func NewMerkleTree(contents []string) *MerkleTree {
	leafs := convert2Node(contents)
	mkl := &MerkleTree{
		leafs: leafs,
	}
	rootHash, rootNode := createMerkleHash(leafs) //生成梅克尔根
	mkl.Root = rootNode
	mkl.RootHash = rootHash
	return mkl
}

//生成梅克尔根
func createMerkleHash(leafs []*MerkleNode) (rootHash []byte, merkleNode *MerkleNode) {
	if len(leafs) == 1 {
		rootHash = leafs[0].hash
		merkleNode = leafs[0]
		return
	}
	var nextFloorNodes []*MerkleNode
	count := 0
	if len(leafs)%2 == 1 {
		count = len(leafs) - 1
	} else {
		count = len(leafs)
	}
	for i := 0; i < count; i += 2 {
		left := leafs[i]
		right := leafs[i+1]

		pHash := sha256.New()
		pHash.Write(left.hash)
		pHash.Write(right.hash)
		parent := pHash.Sum(nil)

		node := &MerkleNode{
			hash:  parent[:],
			left:  left,
			right: right,
		}
		left.Parent = node
		right.Parent = node

		nextFloorNodes = append(nextFloorNodes, node)
	}
	if len(leafs)%2 == 1 {
		nextFloorNodes = append(nextFloorNodes, leafs[len(leafs)-1])
	}
	return createMerkleHash(nextFloorNodes)

}

//校验梅克尔根
func (m *MerkleTree) verify(contents []string) bool {
	leafs := convert2Node(contents)
	rootHash, _ := createMerkleHash(leafs) //生成梅克尔根
	if bytes.Equal(m.RootHash, rootHash) {
		return true
	} else {
		return false
	}
}
func getHash(ori []byte) []byte {
	sha := sha256.Sum256(ori)
	return sha[:]
}

func merkle() {

	var contents []string
	for i := 0; i < 12; i++ {
		contents = append(contents, "test")
	}
	mkl := NewMerkleTree(contents)

	fmt.Printf("梅克尔根：%x\n", mkl.RootHash)

	var tests []string
	for i := 0; i < 12; i++ {
		tests = append(tests, "test")
	}
	fmt.Println("校验结果：", mkl.verify(tests))

}
