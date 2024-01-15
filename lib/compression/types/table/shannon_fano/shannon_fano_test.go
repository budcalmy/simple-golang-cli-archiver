package shannonfano

import (
	"reflect"
	"testing"
)

func Test_findDividePlace(t *testing.T) {
	tests := []struct {
		name  string
		codes []code
		want  int
	}{
		{
			name: "one number",
			codes: []code{{
				Quantity: 2,
			}},
			want: 0,
		},
		{
			name:  "two equal numbers",
			codes: []code{{Quantity: 2}, {Quantity: 2}},
			want:  1,
		},
		{
			name:  "tree numbers",
			codes: []code{{Quantity: 2}, {Quantity: 1}, {Quantity: 1}},
			want:  1,
		},
		{
			name: "many numbers",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1}},
			want: 2,
		},
		{
			name: "uncertainly pos (need right most)",
			codes: []code{
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1}},
			want: 1,
		},
		{
			name: "uncertainly pos (need right most)",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDividerPlace(tt.codes); got != tt.want {
				t.Errorf("findDividePlace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_assignCodes(t *testing.T) {
	tests := []struct {
		name  string
		codes []code
		want  []code
	}{
		{
			name:  "two elements",
			codes: []code{{Quantity: 2}, {Quantity: 2}},
			want: []code{
				{Quantity: 2, Bits: 0, Size: 1},
				{Quantity: 2, Bits: 1, Size: 1}},
		},
		{
			name: "tree elements (certain position)",
			codes: []code{
				{Quantity: 2},  //0
				{Quantity: 1},  //10
				{Quantity: 1}}, //11
			want: []code{
				{Quantity: 2, Bits: 0, Size: 1},
				{Quantity: 1, Bits: 2, Size: 2},
				{Quantity: 1, Bits: 3, Size: 2}},
		},
		{
			name: "tree elements (uncertain position)",
			codes: []code{
				{Quantity: 1},  //0
				{Quantity: 1},  //10
				{Quantity: 1}}, //11
			want: []code{
				{Quantity: 1, Bits: 0, Size: 1},
				{Quantity: 1, Bits: 2, Size: 2},
				{Quantity: 1, Bits: 3, Size: 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assignCodes(tt.codes)

			if !reflect.DeepEqual(tt.codes, tt.want) {
				t.Errorf("got %v, want %v", tt.codes, tt.want)
			}
		})
	}
}

func Test_buildEnTable(t *testing.T) {
	tests := []struct {
		name string
		text string
		want encodingTable
	}{
		{
			name: "base test",
			text: "abbbcc",
			want: encodingTable{
				'a': code{
					Char:     'a',
					Quantity: 1,
					Bits:     3,
					Size:     2,
				},
				'b': code{
					Char:     'b',
					Quantity: 3,
					Bits:     0,
					Size:     1,
				},
				'c': code{
					Char:     'c',
					Quantity: 2,
					Bits:     2,
					Size:     2,
				},
			},
		},
		{
			name: "equal all test",
			text: "aabbcc",
			want: encodingTable{
				'a': code{
					Char:     'a',
					Quantity: 2,
					Bits:     0,
					Size:     1,
				},
				'b': code{
					Char:     'b',
					Quantity: 2,
					Bits:     2,
					Size:     2,
				},
				'c': code{
					Char:     'c',
					Quantity: 2,
					Bits:     3,
					Size:     2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildEnTable(createCharStat(tt.text)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildEnTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
