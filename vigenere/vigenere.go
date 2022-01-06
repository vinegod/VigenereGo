package vigenere

import (
	"bytes"
)

type Vigenere struct {
	Key []byte
}

func (v *Vigenere) code(s []byte, mode int) (n int, text []byte) {
	str := PrepapeString(s)
	keyLen := len(v.Key)

	for i, char := range str {
		num := int(char-'a') + mode*(int(v.Key[i%keyLen])-'a')
		num = 97 + num%26
		if num < 97 {
			num += 26
		}
		str[i] = byte(num)
	}
	return len(str), str
}

func (v *Vigenere) Encode(s []byte) (n int, text []byte) {
	return v.code(s, 1)
}

func (v *Vigenere) Decode(s []byte) (n int, text []byte) {
	return v.code(s, -1)
}

func PrepapeString(s []byte) []byte {
	for _, c := range []byte{' ', ',', '.', '?', '!', '-'} {
		s = bytes.ReplaceAll(s, []byte{c}, []byte{})
	}
	return bytes.ToLower(s)
}
