// aceita um conjunto de dados e o reduz a um fixo menor.
// hashes são usados em tudo em programação, princ. p/ buscar dados e detectar
// em go sao divididos Criptografados e Não criptografados
// lista das Não Crip: adler32, crc32, crc64
// nesse codigo vamos criar uma Hash e escrever nossos dados nele

package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	// cria a hash
	h := crc32.NewIEEE()
	// escreve os dados no hash
	h.Write([]byte("codigo com pacote hash"))
	//calcula o hash
	v := h.Sum32()
	fmt.Println(v)
}