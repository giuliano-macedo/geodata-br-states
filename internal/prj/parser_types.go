package prj

import (
	"fmt"
)

type Node struct {
	Name         string
	Type         NodeType
	LiteralValue float64

	Children []*Node
}

type NodeType int

const (
	NodeUnknown NodeType = iota
	NodeLiteral
	NodeProjcs
	NodeGeogcs
	NodeSpheroid
	NodeProjection
	NodeDatum
	NodeParameter
	NodeUnit
	NodePrimem
)

var nodeTypeTable = [...]string{
	NodeUnknown:    "",
	NodeLiteral:    "",
	NodeProjcs:     "PROJCS",
	NodeGeogcs:     "GEOGCS",
	NodeSpheroid:   "SPHEROID",
	NodeProjection: "PROJECTION",
	NodeDatum:      "DATUM",
	NodeParameter:  "PARAMETER",
	NodeUnit:       "UNIT",
	NodePrimem:     "PRIMEM",
}

func ParseNodeType(value string) (NodeType, error) {
	if value == "" {
		return NodeUnknown, fmt.Errorf("can't parse empty string")
	}

	for i, nValue := range nodeTypeTable {
		if nValue == value {
			return NodeType(i), nil
		}
	}

	return NodeUnknown, fmt.Errorf("couldn't parse nodetype %v", value)
}

func (nt NodeType) String() string {
	return nodeTypeTable[nt]
}
