package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"scs/EAS2"
	"scs/InterfaceA"
	"time"
)

func main() {
	var item InterfaceA.InterfaceOnInter
	itemA := &InterfaceA.InterfaceOnIn{}
	/*item
		*/
	item = itemA
	item.QueryItemName()
}

func aespassword() {
	key := []byte("example key 1234")       //秘钥长度需要时AES-128(16bytes)或者AES-256(32bytes)
	plaintext := []byte("exampleplaintext") //原文必须填充至blocksize的整数倍，填充方法可以参见https://tools.ietf.org/html/rfc5246#section-6.2.3.2

	if len(plaintext)%aes.BlockSize != 0 { //块大小在aes.BlockSize中定义
		panic("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key) //生成加密用的block
	if err != nil {
		panic(err)
	}

	// 对IV有随机性要求，但没有保密性要求，所以常见的做法是将IV包含在加密文本当中
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	fmt.Printf("%x\n", ciphertext)
	//随机一个block大小作为IV
	//采用不同的IV时相同的秘钥将会产生不同的密文，可以理解为一次加密的session
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// 谨记密文需要认证(i.e. by using crypto/hmac)

	fmt.Printf("%x\n", ciphertext)
}

//随机数字
func ReadFull() {
	ciphertext := make([]byte, 10)
	fmt.Println(ciphertext)
	iv := ciphertext[:5]
	fmt.Println(iv)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	fmt.Println(iv)
	fmt.Println(rand.Reader)

}

func password() {
	fmt.Println(time.Now().Unix())
	usera, err := EAS2.Encrypt("d8000180fea811e8869cdd3e281acac2 1554682516 app")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("usera", usera)
	//Encrypt("09cb38a0c2e611e88641a948e4eaedd3")
	fmt.Println("进入进来")
}
