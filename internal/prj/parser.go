package prj

import (
	"fmt"
	"strconv"
	"strings"
)

func paramsShallowSplit(data string) (ans []string) {
	start := 0
	stack := 0
	for i, c := range data {
		switch c {
		case '[':
			stack++
		case ']':
			stack--
		}

		if stack == 0 && c == ',' {
			ans = append(ans, data[start:i])
			start = i + 1
		}
	}

	ans = append(ans, data[start:])
	return
}

func parseTree(data string) (node *Node, err error) {
	node = &Node{}
	openIndex := strings.Index(data, "[")
	if openIndex == -1 {
		node.Type = NodeLiteral
		node.LiteralValue, err = strconv.ParseFloat(data, 64)

		return
	}

	nodeTypeStr := data[0:openIndex]

	node.Type, err = ParseNodeType(nodeTypeStr)
	if err != nil {
		return
	}

	closeIndex := strings.LastIndex(data, "]")

	if closeIndex == -1 {
		return nil, fmt.Errorf("no closing ] for type %v", nodeTypeStr)
	}

	childrenSplit := paramsShallowSplit(data[openIndex+1 : closeIndex])
	if len(childrenSplit) == 0 {
		return nil, fmt.Errorf("empty arguments for type %v", nodeTypeStr)
	}

	node.Name = strings.Trim(childrenSplit[0], "\"")

	node.Children = make([]*Node, len(childrenSplit)-1)
	for i, childrenData := range childrenSplit[1:] {
		if node.Children[i], err = parseTree(childrenData); err != nil {
			return
		}
	}

	return
}
