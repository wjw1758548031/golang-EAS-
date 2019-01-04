package EAS2

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// 需要定义key和偏移量
const (
	key = "MZ7^5f$bjRP#mL$Q"
	iv  = "*JyxSM^Svfzp7wd&"
)

func Encrypt(encodeStr string) (cipherText string, err error) {
	//生成加密用的block
	ckey, err := aes.NewCipher([]byte(key))
	if nil != err {
		fmt.Println("钥匙创建错误:", err)
		return "", err
	}

	//块大小
	blockSize := ckey.BlockSize()

	//转换为[]byte类型
	str := []byte(encodeStr)

	//偏移量转换类型
	iv := []byte(iv)

	fmt.Println("加密的字符串", string(str), "\n加密钥匙", key, "\n向量IV", string(iv))

	fmt.Println("加密前的字节：", str, "\n")

	encrypter := cipher.NewCBCEncrypter(ckey, iv)

	// PKCS7补码
	str = PKCS7Padding(str, blockSize)
	out := make([]byte, len(str))

	encrypter.CryptBlocks(out, str)
	fmt.Println("加密后字节：", out)
	fmt.Println("str：", str)

	// hex 兼容nodejs cropty-js包
	cipherText = hex.EncodeToString(out)
	return cipherText, nil
}

func Decrypt(encodeStr string) (origStr string, err error) {

	ckey, err := aes.NewCipher([]byte(key))
	if nil != err {
		fmt.Println("钥匙创建错误:", err)
		return "", err
	}

	base64Str, _ := hex.DecodeString(encodeStr)
	base64Out := base64.URLEncoding.EncodeToString(base64Str)

	fmt.Println("\n开始解码")
	decrypter := cipher.NewCBCDecrypter(ckey, []byte(iv))

	base64In, _ := base64.URLEncoding.DecodeString(base64Out)

	in := make([]byte, len(base64In))

	decrypter.CryptBlocks(in, base64In)

	fmt.Println("解密后的字节：", in)

	// 去除PKCS7补码
	in = UnPKCS7Padding(in)

	fmt.Println("去PKCS7补码：", in)
	fmt.Println("解密：", string(in))
	origStr = string(in)
	return origStr, nil
}

/**
 *	PKCS7补码
 */
func PKCS7Padding(data []byte, blockSize int) []byte {
	//blockSize := 16
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)

}

/**
 *	去除PKCS7的补码
 */
func UnPKCS7Padding(data []byte) []byte {
	length := len(data)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
