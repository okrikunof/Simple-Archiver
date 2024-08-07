package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Encode(str string) string {
	str = prepareText(str)

	chunks := splitByChunks(encodeBin(str), chunkSize)

	return chunks.ToHex().ToString()
}

func Decode(encodedText string) string {
	//hChunks := NewHexChunks(encodedText)
	//
	//bChunks := hChunks.ToBinary()
	//
	//bString := bChunks.Join()

	return ""
}

func (bcs BinaryChunks) Join() string {
	var buf strings.Builder

	for _, bc := range bcs {
		buf.WriteString(string(bc))
	}

	return buf.String()
}

func (hcs HexChunks) ToBinary() BinaryChunks {
	res := make(BinaryChunks, 0, len(hcs))

	for _, chunk := range hcs {
		bChunk := chunk.ToBinary()

		res = append(res, bChunk)
	}

	return res
}

func (hc HexChunk) ToBinary() BinaryChunk {
	num, err := strconv.ParseUint(string(hc), 16, chunkSize)
	if err != nil {
		panic("can't convert hex chunk to binary" + err.Error())
	}

	res := fmt.Sprintf("%08b", num)

	return BinaryChunk(res)

}

func (hcs HexChunks) ToString() string {
	switch len(hcs) {
	case 0:
		return ""
	case 1:
		return string(hcs[0])
	}

	var buf strings.Builder

	buf.WriteString(string(hcs[0]))

	for _, hc := range hcs[1:] {
		buf.WriteString(hexChunksSeparator)
		buf.WriteString(string(hc))
	}

	return buf.String()
}

func (bcs BinaryChunks) ToHex() HexChunks {
	res := make(HexChunks, 0, len(bcs))

	for _, chunk := range bcs {
		hexChunk := chunk.ToHex()

		res = append(res, hexChunk)
	}

	return res
}

func (bc BinaryChunk) ToHex() HexChunk {
	num, err := strconv.ParseUint(string(bc), 2, chunkSize)
	if err != nil {
		panic("can't convert binary chunks to hex chunks" + err.Error())
	}

	res := strings.ToUpper(fmt.Sprintf("%x", num))

	if len(res) == 1 {
		res = "0" + res
	}

	return HexChunk(res)
}

func encodeBin(str string) string {
	var buf strings.Builder
	for _, ch := range str {
		buf.WriteString(bin(ch))
	}

	return buf.String()
}

func bin(ch rune) string {
	table := getEncodingTable()

	res, ok := table[ch]
	if !ok {
		panic("unknown character" + string(ch))
	}

	return res
}

func getEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		'e': "101",
		't': "1001",
		'o': "10001",
		'n': "10000",
		'a': "011",
		's': "0101",
		'i': "01001",
		'r': "01000",
		'h': "0011",
		'd': "00101",
		'l': "001001",
		'!': "001000",
		'u': "00011",
		'c': "000101",
		'f': "000100",
		'm': "000011",
		'p': "0000101",
		'g': "0000100",
		'w': "0000011",
		'b': "0000010",
		'y': "0000001",
		'v': "00000001",
		'j': "000000001",
		'k': "0000000001",
		'x': "00000000001",
		'q': "000000000001",
		'z': "000000000000",
	}
}

func prepareText(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		if unicode.IsUpper(ch) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}
