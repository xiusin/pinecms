package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
)

func AesEncryptCBC(origData []byte, key []byte) string {
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	key = Md5(key)[:24]
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()               // 获取秘钥块的长度
	origData = pkcs7Padding(origData, blockSize) // 补全码
	//fmt.Println(blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式
	//fmt.Println(string(Md5(key)[:blockSize]))
	encrypted := make([]byte, len(origData))   // 创建数组
	blockMode.CryptBlocks(encrypted, origData) // 加密
	return base64.StdEncoding.EncodeToString(encrypted)
}
func AesDecryptCBC(enc string, key []byte) (string, error) {
	key = Md5(key)[:24]
	encrypted, _ := base64.StdEncoding.DecodeString(enc)
	block, _ := aes.NewCipher(key)                              // 分组秘钥
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式
	decrypted := make([]byte, len(encrypted))                   // 创建数组
	blockMode.CryptBlocks(decrypted, encrypted)                 // 解密
	decrypted, err := pkcs7UnPadding(decrypted)                 // 去除补全码
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs7UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	unpadding := int(origData[length-1])
	if length-unpadding < 0 {
		return []byte{}, errors.New("解密失败")
	}
	return origData[:(length - unpadding)], nil
}

func Md5(str []byte) []byte {
	h := md5.New()
	h.Write(str)
	//fmt.Println(hex.EncodeToString(h.Sum(nil)))
	return []byte(hex.EncodeToString(h.Sum(nil)))
}
