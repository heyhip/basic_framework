package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"general_server/core/log"
)

func Sha1(str string) (sha1String string) {
	// 创建散列
	s := sha1.New()

	// 写入要处理的字节
	s.Write([]byte(str))

	// 最终的散列值的字符切片
	bs := s.Sum(nil)

	// SHA1 值经常以 16 进制输出，使用%x 来将散列结果格式化为 16 进制字符串
	return fmt.Sprintf("%x", bs)
}

func Md5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	bs := s.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

//加密数据
func EncodeAes(data, key, iv []byte) []byte {
	aesBlockEncrypter, err := aes.NewCipher(key)
	content := PKCS5Padding(data, aesBlockEncrypter.BlockSize())
	encrypted := make([]byte, len(content))
	if err != nil {
		log.Error(err)
		return nil
	}
	aesEncrypter := cipher.NewCBCEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.CryptBlocks(encrypted, content)
	return encrypted
}

//解密数据
func DecodeAes(src, key, iv []byte) []byte {
	var err error
	decrypted := make([]byte, len(src))
	var aesBlockDecrypter cipher.Block
	aesBlockDecrypter, err = aes.NewCipher(key)
	if err != nil {
		log.Error(err)
		return nil
	}
	aesDecrypter := cipher.NewCBCDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.CryptBlocks(decrypted, src)
	return PKCS5Trimming(decrypted)
}

/**
PKCS5包装
*/
func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

/*
解包装
*/
func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
