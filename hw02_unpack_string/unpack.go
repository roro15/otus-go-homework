package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

const escapeRune = '\\'

type RuneReader struct {
	input []rune
	index int
}

func NewRuneReader(input []rune) *RuneReader {
	return &RuneReader{
		input: input,
	}
}

func (r *RuneReader) Read() (rune, bool) {
	if r.index < len(r.input) {
		v := r.input[r.index]
		r.index++
		return v, true
	}
	return 0, false
}

func Unpack(input string) (string, error) {
	reader := NewRuneReader([]rune(input))
	builder := strings.Builder{}
	var current, next rune
	var ok bool
	current, ok = reader.Read()
	if !ok {
		return "", nil
	}
	if unicode.IsDigit(current) {
		return "", ErrInvalidString
	}
	for {
		if current == escapeRune {
			next, ok = reader.Read()
			if !ok || !(next == escapeRune || unicode.IsDigit(next)) {
				return "", ErrInvalidString
			}
			current = next
		}
		next, ok = reader.Read()
		if !ok {
			builder.WriteRune(current)
			break
		}
		if unicode.IsDigit(next) {
			n := int(next - '0')
			builder.WriteString(strings.Repeat(string(current), n))
			current, ok = reader.Read()
			if !ok {
				break
			}
			if unicode.IsDigit(current) {
				return "", ErrInvalidString
			}
			continue
		}
		builder.WriteRune(current)
		current = next
	}
	return builder.String(), nil
}
