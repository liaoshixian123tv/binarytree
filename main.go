package main

import (
	"binarytree/model/node"
	"encoding/json"
	"fmt"
)

func main() {

	tmp := `[[{"name":"A","parent":""}],[{"name":"B","parent":"A"},{"name":"C","parent":"A"}],[{"name":"D","parent":"B"},{"name":"E","parent":"C"},{"name":"E","parent":"C"}],[{"name":"G","parent":"D"},{"name":"H","parent":"D"},{"name":"I","parent":"F"}]]`
	var input [][]node.Input
	if err := json.Unmarshal([]byte(tmp), &input); err != nil {
		fmt.Println(err)
	}
	// var tree *node.Node

}
