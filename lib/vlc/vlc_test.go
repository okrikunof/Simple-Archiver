package vlc

import (
	"reflect"
	"testing"
)

func Test_encodeBin(t *testing.T) {
	tests := []struct {
		name string
		bStr string
		want string
	}{
		{
			name: "base test",
			bStr: "!xyz",
			want: "001000000000000010000001000000000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeBin(tt.bStr); got != tt.want {
				t.Errorf("encodeBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunks_ToHex(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want HexChunks
	}{
		{
			name: "base test",
			bcs:  BinaryChunks{"0101111", "10000000"},
			want: HexChunks{"2F", "80"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.ToHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncoe1(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "My name is Ted",
			want: "20 30 3C 18 77 4A E4 4D 28",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.str); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prepareText(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "My name is Ted",
			want: "!my name is !ted",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareText(tt.str); got != tt.want {
				t.Errorf("prepareText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunk_ToBinary(t *testing.T) {
	tests := []struct {
		name string
		hc   HexChunk
		want BinaryChunk
	}{
		{
			name: "base test",
			hc:   HexChunk("2F"),
			want: BinaryChunk("00101111"),
		},
		{
			name: "base test",
			hc:   HexChunk("80"),
			want: BinaryChunk("10000000"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hc.ToBinary(); got != tt.want {
				t.Errorf("ToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunks_ToBinary(t *testing.T) {
	tests := []struct {
		name string
		hcs  HexChunks
		want BinaryChunks
	}{
		{
			name: "base test",
			hcs:  HexChunks{"2F", "80"},
			want: BinaryChunks{"00101111", "10000000"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hcs.ToBinary(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBinary() = %v, want %v", got, tt.want)
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
			bcs:  BinaryChunks{"0101111", "10000000"},
			want: "010111110000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.Join(); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name       string
		encodeText string
		want       string
	}{
		{
			name:       "base test",
			encodeText: "20 30 3C 18 77 4A E4 4D 28",
			want:       "My name is Ted",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decode(tt.encodeText); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
