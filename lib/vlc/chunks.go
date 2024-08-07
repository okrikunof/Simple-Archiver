package vlc

import (
	"strings"
	"unicode/utf8"
)

type BinaryChunk string

type BinaryChunks []BinaryChunk

type HexChunk string

type HexChunks []HexChunk

type encodingTable map[rune]string

const chunkSize = 8

const hexChunksSeparator = " "

func NewHexChunks(str string) HexChunks {
	parts := strings.Split(str, hexChunksSeparator)

	res := make(HexChunks, 0, len(parts))

	for _, part := range parts {
		res = append(res, HexChunk(part))
	}

	return res
}

func splitByChunks(bStr string, chunkSize int) BinaryChunks {
	strLen := utf8.RuneCountInString(bStr)

	chunksCount := strLen / chunkSize

	if strLen/chunkSize != 0 {
		chunksCount++
	}

	res := make(BinaryChunks, 0, chunksCount)

	var buf strings.Builder

	for i, ch := range bStr {
		buf.WriteString(string(ch))

		if (i+1)%chunkSize == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}

	if buf.Len() != 0 {
		lastChunk := buf.String()

		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))

		res = append(res, BinaryChunk(lastChunk))
	}

	return res
}
