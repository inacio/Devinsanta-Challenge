//Esta solução não está com bom desempenho, gostaria que ajudassem a melhora-la se possivel

package main

import (
	"fmt"
	"log"
	"math/big"
	"time"
)

func factorial(n *big.Int) (result *big.Int) {
	//fmt.Println("n = ", n)
	result = new(big.Int)

	switch n.Cmp(&big.Int{}) {
	case -1, 0:
		result.SetInt64(1)
	default:
		result.Set(n)
		var one big.Int
		one.SetInt64(1)
		result.Mul(result, factorial(n.Sub(n, &one)))
	}
	return
}

func main() {
	start := time.Now()
	limite := int64(299)
	array := make([]*big.Int, limite)

	cont := int64(2)
	i := int64(0)
	n := big.NewInt(2)
	for i < limite {
		fmt.Print(n, " ")
		//go func() { //travou meu computador ao utilizar o goroutines, sei que iria embaralhar mas gostaria de ver como ficaria o desempenho (usou 100% de todos os meus núcleos) :(
		array[i] = factorial(n.Mul(n, n))
		cont++
		n = big.NewInt(int64(cont))
		i++
		//}()
	}

	fmt.Println(array)

	elapsed := time.Since(start)
	log.Printf("\n\n\n Tempo para executar: %s", elapsed)
}
