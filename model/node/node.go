package node

// Input 輸入
type Input struct {
	Name   string `json:"name"`
	Parent string `json:"parent"`
}

// Node 節點
type Node struct {
	Level    int    `json:"level"`
	Name     string `json:"name"`
	Children []Node `json:"children"`
}

// NewNodeByName 新增節點
func (c *Input) NewNodeByName(level int, child []Node) Node {
	return Node{Name: c.Name, Level: level, Children: child}
}

// NewNodeByParent 新增節點
func (c *Input) NewNodeByParent(level int, child []Node) Node {
	return Node{Name: c.Parent, Level: level, Children: child}
}

// AddChild 新增Child
func (c *Node) AddChild(args ...Node) {
	c.Children = append(c.Children, args...)
}
