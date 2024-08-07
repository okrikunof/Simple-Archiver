package vlc

import (
	"reflect"
	"testing"
)

func Test_splitByChunks(t *testing.T) {
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
			name: "base test",
			args: args{
				bStr:      "001000000000000010000001000000000000",
				chunkSize: 8,
			},
			want: BinaryChunks{"00100000", "00000000", "10000001", "00000000", "00000000"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByChunks(tt.args.bStr, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHexChunks(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want HexChunks
	}{
		{
			name: "base test",
			str:  "001000000000000010000001000000000000",
			want: HexChunks{"00100000", "00000000", "10000001", "00000000", "00000000"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHexChunks(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHexChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}
