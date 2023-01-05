package main

import (
	"binarytree/model/node"
	"encoding/json"
	"fmt"
)

// jsonFormat json格式的題目
var jsonFormat string

// input 題目
var input [][]node.Input

func init() {
	jsonFormat = `[[{"name":"A","parent":""}],[{"name":"B","parent":"A"},{"name":"C","parent":"A"}],[{"name":"D","parent":"B"},{"name":"E","parent":"C"},{"name":"F","parent":"C"}],[{"name":"G","parent":"D"},{"name":"H","parent":"D"},{"name":"I","parent":"F"}]]`
	if err := json.Unmarshal([]byte(jsonFormat), &input); err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	var tree node.Node
	nodeMap := make(map[string]node.Node)
	for i := len(input) - 1; i >= 0; i-- {
		nodeInput := input[i]
		for _, info := range nodeInput {
			if botVal, botOK := nodeMap[info.Name]; !botOK { // 檢查下面的node是否存在
				// 不存在 則創建
				tmpBotVal := info.NewNodeByName(i, []node.Node{})
				nodeMap[info.Name] = tmpBotVal
				if info.Parent == "" {
					continue
				}
				var updateTopVal node.Node
				// 檢查所屬節點是否存在
				if topVal, topOK := nodeMap[info.Parent]; !topOK {
					// 不存在 則創建
					updateTopVal = info.NewNodeByParent(i-1, []node.Node{tmpBotVal})
				} else {
					// 存在 則加入該node
					topVal.AddChild(tmpBotVal)
					updateTopVal = topVal
				}
				nodeMap[info.Parent] = updateTopVal
			} else {
				if info.Parent == "" { // 若無所屬節點 則視為開始節點
					tree = botVal
					continue
				}
				var updateTopVal node.Node
				if topVal, topOK := nodeMap[info.Parent]; !topOK { // 檢查所屬節點是否存在
					// 不存在 則創建
					updateTopVal = info.NewNodeByParent(i-1, []node.Node{botVal})
				} else {
					// 存在 則加入該node
					topVal.AddChild(botVal)
					updateTopVal = topVal
				}
				nodeMap[info.Parent] = updateTopVal
			}
		}
	}

	readAllNodes(tree)
}

// readAllNodes 遍利所有節點
func readAllNodes(tmpnode node.Node) {
	fmt.Printf("%+v \n", tmpnode)
	if len(tmpnode.Children) > 0 {
		for _, child := range tmpnode.Children {
			readAllNodes(child)
		}
	}
}
