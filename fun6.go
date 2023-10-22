package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	strs := []string{"Hello", " ", "World", "!"}

	media := calcularMedia(nums)
	fmt.Println("Média:", media)

	vogais := contarVogais("Hello World!")
	fmt.Println("Vogais:", vogais)

	concatenado := concatenarStrings(strs)
	fmt.Println("Concatenado:", concatenado)

	segundoMaior := encontrarSegundoMaior(nums)
	fmt.Println("Segundo maior:", segundoMaior)

	posicao := encontrarPosicao(nums, 3)
	fmt.Println("Posição:", posicao)
}
