package haffman

import (
	"archiver/lib/compression/types/table"
	"sort"
)

type Generator struct{}

type charStat map[rune]int

func NewGenerator() Generator {
	return Generator{}
}

type encodingTable map[rune]string

type haffmanNode struct {
	Char    rune
	Freq    int
	Left    *haffmanNode
	Right   *haffmanNode
	Bits    uint32
	BitSize int
}

func (g Generator) NewTable(text string) table.EncodingTable {

	res := make(encodingTable)

	charStat := createCharProbs(text)

	root := buildHaffmanTree(charStat)

	generateCharCodes(root, "", res)

	return table.EncodingTable(res)
}

func generateCharCodes(root *haffmanNode, code string, codes map[rune]string) {
	if root.Left == nil && root.Right == nil {
		codes[root.Char] = code
		return
	}

	generateCharCodes(root.Left, code+"0", codes)
	generateCharCodes(root.Right, code+"1", codes)
}

func buildHaffmanTree(cp charStat) *haffmanNode {

	var nodes []*haffmanNode

	for ch, freq := range cp {
		nodes = append(nodes, &haffmanNode{
			Char: ch,
			Freq: freq,
		})
	}

	for len(nodes) > 1 {
		sort.Slice(nodes, func(i, j int) bool {
			return nodes[i].Freq < nodes[j].Freq
		})

		left := nodes[0]
		right := nodes[1]

		merged := &haffmanNode{
			Freq:  left.Freq + right.Freq,
			Left:  left,
			Right: right,
		}

		nodes = nodes[2:]
		nodes = append(nodes, merged)
	}

	return nodes[0]
}

func createCharProbs(text string) charStat {
	res := make(charStat)

	total := 0

	for _, ch := range text {
		res[ch]++
		total++
	}

	return res
}
