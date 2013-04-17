package gouri

import (
	"crypto/rand"
	"io"
)

var (
	chars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
)

func New(len int) string {
	return (RandomString(len, chars))
}

func RandomString(length int, char_pool []byte) string {
	randstring := make([]byte, length)
	randbyte := make([]byte, length+(length/4))
	clen := byte(len(chars))

	for i, _ := range randstring {
		if _, err := io.ReadFull(rand.Reader, randbyte); err != nil {
			panic("error : " + err.Error())
		}
		for _, r := range randbyte {
			randstring[i] = char_pool[r%clen]
		}

	}
	return string(randstring)
}
