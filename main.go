package main

import (
	"fmt"

	"github.com/SocioExml/godesde0/ejercicios"
)

func main() {
	//variables.RestoVariables()

	/*estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)*/

	/*if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es Windows")
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "Otro":
		fmt.Println("Esto es otro")
	default:
		fmt.Printf("%s \n", os)
	}*/

	numero, texto := ejercicios.ConvNumerico("500")
	fmt.Println(numero)
	fmt.Println(texto)
}
