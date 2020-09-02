package main

import (
	"fmt"

	"github.com/PIPE1303/First/calculadora"
)

func main() {
	var saludo = "Estoy saludando"
	fmt.Printf("Hola mundo %v. \nVariable del tipo: %T\n", saludo, saludo)

	var numero1 float32 = 12.3
	var numero2 float32 = 0.0

	fmt.Printf("Suma: %v\n", calculadora.Sumar(numero1, numero2))
	fmt.Printf("Resta: %v\n", calculadora.Restar(numero1, numero2))
	fmt.Printf("Multiplicación: %v\n", calculadora.Multiplicar(numero1, numero2))

	var res, err = calculadora.Dividir(numero1, numero2)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("División: %d\n", res)
	}
}
