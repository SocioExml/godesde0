package main

import (
	g "github.com/SocioExml/godesde0/middleware"
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

	/*numero, texto := ejercicios.ConvNumerico("500")
	fmt.Println(numero)
	fmt.Println(texto)*/

	//teclado.IngresoNumeros()

	//iteraciones.Iterar()
	//fmt.Println(ejercicios.TabladMultiplicar())
	//files.GrabaTabla()

	//files.SumaTabla()

	//funciones.Calculos()

	//funciones.LlamarClosure()

	//funciones.Exponencia(2)

	//arreglos_slices.MuestroArreglos()

	//arreglos_slices.Capacidad()

	// mapas.MostrarMapas()
	//	users.AltaUsuario()

	/*Pedro := new(modelos.Hombre)
	e.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)*/

	/*
		canal1 := make(chan bool)
		go g.MiNombreLentoooo("Edgar Molina", canal1)
		defer func() {
			<-canal1
		}()
		fmt.Println("Estoy aqui")
	*/

	g.MiMiddleware()
}
