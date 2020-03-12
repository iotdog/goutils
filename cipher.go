package goutils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"

	"github.com/leesper/holmes"
)

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

//AesEcbPkcs5padding .
func AesEcbPkcs5padding(plainData, pwd []byte) []byte {
	block, err := aes.NewCipher(pwd)
	if err != nil {
		holmes.Infoln("key error", err)
		return nil
	}
	ecb := NewECBEncrypter(block)
	plainData = pkcs5Padding(plainData, block.BlockSize())
	// holmes.Infoln(hex.EncodeToString(plainData))
	crypted := make([]byte, len(plainData))
	ecb.CryptBlocks(crypted, plainData)
	return crypted
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// AesEcbNoPaddingDec AES ECB NoPadding解密
func AesEcbNoPaddingDec(cipherData, pwd []byte) []byte {
	block, err := aes.NewCipher(pwd)
	if err != nil {
		holmes.Infoln("key error", err)
		return nil
	}
	ecb := NewECBDecrypter(block)
	plainData := make([]byte, len(cipherData))
	ecb.CryptBlocks(plainData, cipherData)
	return plainData
}

type ecbEncrypter ecb

// NewECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

func (x *ecbDecrypter) BlockSize() int { return x.blockSize }

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
