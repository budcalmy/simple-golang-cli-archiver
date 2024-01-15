package haffman

import (
	"archiver/lib/compression/types/table"
	"reflect"
	"testing"
)

func Test_buildHaffmanTree(t *testing.T) {
	tests := []struct {
		name string
		cp   charStat
		want *haffmanNode
	}{
		{
			name: "base test",
			cp: charStat{
				'a': 3,
				'b': 2,
				'c': 6,
			},
			want: &haffmanNode{
				Freq: 11,
				Left: &haffmanNode{
					Freq: 5,
					Left: &haffmanNode{
						Char:  'b',
						Freq:  2,
						Left:  nil,
						Right: nil,
					},
					Right: &haffmanNode{
						Char:  'a',
						Freq:  3,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &haffmanNode{
					Char:  'c',
					Freq:  6,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildHaffmanTree(tt.cp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildHaffmanTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerator_NewTable(t *testing.T) {
	tests := []struct {
		name string
		g    Generator
		text string
		want table.EncodingTable
	}{
		{
			name: "base test",
			g:    Generator{},
			text: "aaabbcccccddddddddeeeeee",
			want: table.EncodingTable{
				'a': "001",
				'b': "000",
				'c': "01",
				'd': "11",
				'e': "10",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Generator{}
			if got := g.NewTable(tt.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generator.NewTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
