package leetcode

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/description/

type Codec struct {
}

func CodecConstructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *IntTreeNode) string {
	var tokens []string

	var dfs func(node *IntTreeNode)
	dfs = func(node *IntTreeNode) {
		if node == nil {
			tokens = append(tokens, "null")
			return
		}

		val := strconv.Itoa(node.Val)
		tokens = append(tokens, val)

		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	serialized := strings.Join(tokens, ",")
	return serialized
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *IntTreeNode {
	tokens := strings.Split(data, ",")

	var (
		i   int
		dfs func() *IntTreeNode
	)

	dfs = func() *IntTreeNode {
		if tokens[i] == "null" {
			i++
			return nil
		}

		val, _ := strconv.Atoi(tokens[i])
		i++

		node := IntTreeNode{
			Val: val,
		}

		node.Left = dfs()
		node.Right = dfs()

		return &node
	}

	tree := dfs()
	return tree
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
