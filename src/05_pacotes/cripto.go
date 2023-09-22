// A cripto são os exemplos criptografados
//  vamos usar uma cmuito comum a SHA-1

package main

import (
	"fmt"
	"crypto/sha1"
)

func main() {
	h := sha1.New()
	h.Write([]byte("codigo com pacote hash"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}