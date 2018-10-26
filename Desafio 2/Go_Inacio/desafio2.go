package main

import (
	"fmt"
	"time"
	"log"
)

func main() {
	start := time.Now()
  	size:= 300
    array := make([]int, size-1)
    n := 2;
    i := 0;
    for n <= size {
    	array[i] = n * n
    	n++
    	i++
    }
    
    fmt.Println(array)
    
    elapsed := time.Since(start)
    log.Printf("\n\n\n Tempo para executar: %s", elapsed)
}
