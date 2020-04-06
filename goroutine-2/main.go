package main
import "fmt"
func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}
func main() {
	// Suponha que temos uma função de chamada f(s).
	// Aqui está como nós chamamos ela da maneira usual, executando sincronicamente
    f("direto")

	// Para chamar essa função em uma goroutine, use go f(s).
	// Essa nova goroutine executará simultaneamente com a que foi chamada
    go f("goroutine")

	// Você pode também iniciar uma goroutine para uma função de chamada anônima.
    go func(msg string) {
        fmt.Println(msg)
    }("indo")

	// Nossas duas goroutines são executadas de forma assíncrona em goroutines separadas agora,
	// assim a goroutine cai até aqui. Esse código Scanln exige que pressionamos uma tecla antes de sair do programa
    var input string
    fmt.Scanln(&input)
    fmt.Println("pronto")
}