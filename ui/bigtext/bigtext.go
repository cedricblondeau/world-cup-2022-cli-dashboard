package bigtext

import (
	"fmt"
	"strings"
)

var (
	chars      = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "-", "?"}
	charHeight = 11
)

type BigText struct {
	bigTextByChar map[string]string
}

func NewBigText() *BigText {
	lines := strings.Split(font, "\n")

	if len(lines) != len(chars)*charHeight {
		panic(fmt.Errorf("BigTextFont: want %d lines in font file, got %d", len(chars)*charHeight, len(lines)))
	}

	bigTextByChar := make(map[string]string, len(chars))
	for i, char := range chars {
		startIndex := i * charHeight
		bigTextByChar[char] = strings.Join(lines[startIndex:startIndex+charHeight], "\n")
	}

	return &BigText{bigTextByChar}
}

func (b BigText) Char(char string) string {
	unknown, ok := b.bigTextByChar["?"]
	if !ok {
		panic("BigTextFont: cannot find ? char")
	}

	if bigtext, ok := b.bigTextByChar[char]; ok {
		return bigtext
	}

	return unknown
}
