package main

import "fmt"

func main() {
	n := 3

	// É aqui que se cria o canal, que pode ser usado
	// para mover o tipo 'int'
	out := make(chan int)

	// Ainda rodamos a função como uma goroutine, mas
	// o canal que criamos também é fornecido
	go multiplyByTwo(n, out)

	// Uma vez que quaquer retorno é recebido nesse canal, printa no console e continua
	fmt.Println(<-out)
}

// A funcção agora aceita um canal como segundo parametro
func multiplyByTwo(num int, out chan<- int) {
	result := num * 2

	// ... e faz um pipe do resultado
	out <- result

}