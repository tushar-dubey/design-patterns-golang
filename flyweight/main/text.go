package main

import (
	"strings"
	"unicode"
)

type FormattedText struct {
	text    string
	capital []bool
}

func (f *FormattedText) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(f.text); i++ {
		if f.capital[i] {
			sb.WriteRune(unicode.ToUpper(rune(f.text[i])))
		} else {
			sb.WriteRune(rune(f.text[i]))
		}
	}
	return sb.String()
}

func NewFormattedText(text string) *FormattedText {
	return &FormattedText{text: text, capital: make([]bool, len(text))}
}

func (f *FormattedText) Capitalize(start, end int) {
	for i := start; i <= end; i++ {
		f.capital[i] = true
	}
}

type BetterFormattedText struct {
	text   string
	ranges []*TextRange
}

func (b *BetterFormattedText) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(b.text); i++ {
		c := b.text[i]
		for _, r := range b.ranges {
			if r.InRange(i) && r.Capital {
				c = uint8(unicode.ToUpper(rune(c)))
			}
		}
		sb.WriteRune(rune(c))
	}
	return sb.String()
}

func NewBetterFormattedText(text string) *BetterFormattedText {
	return &BetterFormattedText{text: text}
}

type TextRange struct {
	start, end            int
	Capital, Bold, Italic bool
}

func (b *BetterFormattedText) Range(start, end int) *TextRange {
	tr := &TextRange{
		start:   start,
		end:     end,
		Capital: false,
		Bold:    false,
		Italic:  false,
	}
	b.ranges = append(b.ranges, tr)
	return tr
}

func (tr *TextRange) Capitalize() {
	tr.Capital = true
}

func (tr *TextRange) InRange(index int) bool {
	return tr.start <= index && tr.end >= index
}
