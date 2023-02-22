package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"git.jiaxianghudong.com/go/utils"
)

// DesEncrypt DES-CBC加密, 参数及返回值皆使用字符串
// result[0] base64编码后的密文
// result[1] 错误信息, 无错误返回 ok
func DesEncryptStr(content, key string) (string, string) {
	ret, err := DesEncrypt([]byte(content), []byte(key))
	if err != nil {
		return "", err.Error()
	}
	return utils.Base64Encode(ret), "ok"
}

// DesEncrypt DES-CBC解密, 参数及返回值皆使用字符串
// param: ciphertext 需base64编码
// result[0] 解密结果
// result[1] 错误信息, 无错误返回 ok
func DesDecryptStr(ciphertext, key string) (string, string) {
	ret, err := DesDecrypt(utils.Base64Decode(ciphertext), []byte(key))
	if err != nil {
		return "", err.Error()
	}
	return string(ret), "ok"
}

// DesEncrypt DES-CBC加密
func DesEncrypt(content, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	content = PKCS7Padding(content, blockSize)
	// 向量 (key[:blockSize]) 是密钥的前 blockSize (16) 个字节
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	ciphertext := make([]byte, len(content))
	blockMode.CryptBlocks(ciphertext, content)
	return ciphertext, nil
}

// DesEncrypt DES-CBC解密
func DesDecrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	oriData := make([]byte, len(ciphertext))
	blockMode.CryptBlocks(oriData, ciphertext)
	oriData = PKCS7UnPadding(oriData)
	return oriData, nil
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(oriData []byte) []byte {
	length := len(oriData)
	unpadding := int(oriData[length-1])
	if length < unpadding {
		return []byte{}
	}
	return oriData[:(length - unpadding)]
}
