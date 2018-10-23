package main

import (
	"fmt"
)

func main() {
	a := "Santa"
	b := "rem"
  	cont := 1
    	for cont  <= 200 {
          if cont % 5 == 0 && cont % 6 == 0 {
            fmt.Println(cont,a,b)
          } else if cont  % 6 == 0 {
            fmt.Println(cont ," ",b)
          } else if cont  % 5 == 0 {
            fmt.Println(cont ," ",a)
          }else{
            fmt.Println(cont)
          }

          cont  = cont + 1
    	}
}


//teste: $ go run main.go
