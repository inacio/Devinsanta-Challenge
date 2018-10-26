//Esta solução não está com bom desempenho, gostaria que ajudassem a melhora-la se possivel
package main

import (
	"fmt"
	"log"
	"math/big"
	"time"
)

func main() {
	start := time.Now()
	limite := int64(299)
	array := make([]*big.Int, limite)

	cont := int64(2)
	i := int64(0)
	//n := big.NewInt(2)
	var f big.Int
	for i < limite {
		//fmt.Print(cont, " ")
		//go func() { //
		array[i] = f.MulRange(int64(1), int64(cont*cont))
		//fmt.Println(&f)//apenas com este print 36s para finalizar
		cont++
		//n = big.NewInt(int64(cont))
		i++
		//}()
	}

	fmt.Println(&array)//2.26min para finalizar

	elapsed := time.Since(start)
	log.Printf("\n\n\n Tempo para executar: %s", elapsed)
}
