package vlc

import (
	"reflect"
	"testing"
)

func Test_splitByChunk(t *testing.T) {
	type args struct {
		bStr      string
		chunkSize int
	}
	tests := []struct {
		name string
		args args
		want BinaryChunks
	}{
		{
			name: "append zeroes to chunk",
			args: args{
				bStr:      "10011001100110011001",
				chunkSize: 8,
			},
			want: BinaryChunks{"10011001", "10011001", "10010000"},
		},
		{
			name: "usual test",
			args: args{
				bStr:      "100110011001100110011001",
				chunkSize: 8,
			},
			want: BinaryChunks{"10011001", "10011001", "10011001"},
		},
		{
			name: "empty test",
			args: args{
				bStr:      "",
				chunkSize: 8,
			},
			want: BinaryChunks{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByChunk(tt.args.bStr, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByChunk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunks_Join(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want string
	}{
		{
			name: "base test",
			bcs:  BinaryChunks{"10011101", "00011111", "00111010"},
			want: "100111010001111100111010",
		},
		{
			name: "empty test",
			bcs:  BinaryChunks{},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.Join(); got != tt.want {
				t.Errorf("BinaryChunks.Join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBinChunks(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want BinaryChunks
	}{
		{
			name: "base test", 
			data: []byte{20,30,60,18},
			want: BinaryChunks{"00010100", "00011110", "00111100", "00010010"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBinChunks(tt.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}
