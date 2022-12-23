package node

// Input 輸入
type Input struct {
	Name   string `json:"name"`
	Parent string `json:"parent"`
}

// Node 節點
type Node struct {
	Level    int     `json:"level"`
	Name     string  `json:"name"`
	Children []*Node `json:"children"`
}

// NewNode 新增節點
func NewNode(level int, name string) *Node {
	var newNode *Node
	newNode.Level = level
	newNode.Name = name
	return newNode
}
