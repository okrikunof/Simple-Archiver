package vlc

import (
	"strings"
	"unicode"
)

type EncoderDecoder struct{}

func New() EncoderDecoder {
	return EncoderDecoder{}
}

func (_ EncoderDecoder) Encode(str string) []byte {
	str = prepareText(str)

	chunks := splitByChunks(encodeBin(str), chunkSize)

	return chunks.Bytes()
}

func (_ EncoderDecoder) Decode(encodedData []byte) string {
	bString := NewBinChunks(encodedData).Join()

	dTree := getEncodingTable().DecodingTree()

	return exportText(dTree.Decode(bString))
}

func (bcs BinaryChunks) Join() string {
	var buf strings.Builder

	for _, bc := range bcs {
		buf.WriteString(string(bc))
	}

	return buf.String()
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
		panic("unknown character: " + string(ch))
	}

	return res
}

func getEncodingTable() encodingTable {
	return encodingTable{
		' ':  "111",
		'e':  "1101",
		't':  "11001",
		'o':  "11000",
		'n':  "1011",
		'a':  "10101",
		's':  "10100",
		'i':  "10011",
		'r':  "10010",
		'h':  "10001",
		'd':  "10000",
		'l':  "0111",
		'!':  "01101",
		'u':  "01100",
		'c':  "01011",
		'f':  "01010",
		'm':  "01001",
		'p':  "010001",
		'g':  "010000",
		'w':  "00111",
		'b':  "001101",
		'y':  "001100",
		'v':  "001011",
		'j':  "001010",
		'k':  "001001",
		'x':  "001000",
		'q':  "000111",
		'z':  "000110",
		'\n': "000101",
	}
}

func exportText(str string) string {
	var buf strings.Builder

	var isCapital bool

	for _, ch := range str {
		if isCapital {
			buf.WriteRune(unicode.ToUpper(ch))
			isCapital = false

			continue
		}

		if ch == '!' {
			isCapital = true

			continue
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
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
