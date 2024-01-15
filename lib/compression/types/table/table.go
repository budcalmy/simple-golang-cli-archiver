package table

import "strings"

type decodingTree struct {
	Value    string
	LeftZero *decodingTree
	RightOne *decodingTree
}

type EncodingTable map[rune]string

type Generator interface {
	NewTable(text string) EncodingTable
}

func (et EncodingTable) DecodeForTbl(str string) string {
	dt := et.DecodingTree()

	return dt.DecodeForTree(str)
}

func (enTable EncodingTable) DecodingTree() decodingTree {
	res := decodingTree{}

	for ch, code := range enTable {
		res.add(code, ch)
	}

	return res

}

func (dt *decodingTree) add(code string, value rune) {
	currentNode := dt

	for _, ch := range code {
		switch ch {
		case '0':
			if currentNode.LeftZero == nil {
				currentNode.LeftZero = &decodingTree{}
			}

			currentNode = currentNode.LeftZero
		case '1':
			if currentNode.RightOne == nil {
				currentNode.RightOne = &decodingTree{}

			}
			currentNode = currentNode.RightOne
		}
	}

	currentNode.Value = string(value)
}

func (dt *decodingTree) DecodeForTree(str string) string {
	var buf strings.Builder

	currentNode := dt

	for _, ch := range str {
		if currentNode.Value != "" {
			buf.WriteString(currentNode.Value)
			currentNode = dt
		}

		switch ch {
		case '0':
			currentNode = currentNode.LeftZero
		case '1':
			currentNode = currentNode.RightOne
		}
	}

	if currentNode.Value != "" {
		buf.WriteString(currentNode.Value)
		currentNode = dt
	}

	return buf.String()
}
