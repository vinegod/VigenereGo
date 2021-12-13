package vigenere

import "strings"

const (
	alphabet    = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	alphabetLen = 52
)

type Vigenere struct {
	Text []byte
	Key  []byte
}

func (v *Vigenere) Encode(mode int) (n int, text []byte) {

	textLen, keyLen := v.Len()
	text = make([]byte, textLen)

	v.byteToIndex()

	for i, char := range v.Text {
		text[i] = (alphabetLen + char + byte(1*mode)*v.Key[i%keyLen]) % alphabetLen
	}
	indexToByte(text)

	return textLen, text

}

func (v *Vigenere) Len() (t, k int) {
	t = len(v.Text)
	k = len(v.Key)
	return
}

func byteToIndex(b []byte) {
	for i, char := range b {
		if char == ' ' {
			b[i] = ' '
		} else {
			b[i] = byte(strings.IndexByte(alphabet, char))
		}
	}
}

func indexToByte(b []byte) {
	for i, char := range b {
		if char == ' ' {
			b[i] = ' '
		} else {
			b[i] = alphabet[char]
		}
	}
}

func (v *Vigenere) byteToIndex() {
	byteToIndex(v.Text)
	byteToIndex(v.Key)
}
