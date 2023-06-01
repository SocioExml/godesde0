package middleware

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func MiMiddleware() {
	fmt.Println("Inicio")

	result := operacionMidd(sumar)(2, 3)
	fmt.Println(result)

	result = operacionMidd(restar)(5, 3)
	fmt.Println(result)

	result = operacionMidd(multiplicar)(2, 3)
	fmt.Println(result)

}

func operacionMidd(f func(int, int) int) func(int, int) int {

	return func(a, b int) int {
		fmt.Println("Inicio de Operación")
		return f(a, b)

	}

}
