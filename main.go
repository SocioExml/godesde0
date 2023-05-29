package main

import (
	"fmt"

	"github.com/SocioExml/godesde0/variables"
)

func main() {
	//variables.RestoVariables()

	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)

}
