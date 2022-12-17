//Golang Tutorial for Beginners | Full Go Course
//https://www.youtube.com/watch?v=yyUHQIec83I

package main

import "fmt"

func main() {
	//Atribuicoes e Arrays
	//https://www.digitalocean.com/community/tutorials/understanding-arrays-and-slices-in-go
	//var numbers [3]int //Se declarar precisar implementar!
	palavras := [4]string{"blue coral", "staghorn coral", "pillar coral", "elkhorn coral"}

	//fmt.Println(palavras[1])
	fmt.Println(palavras[1]) //Printando um slice (indice do array, mas com conceito mutavel)

	//Concatenacoes
	var conferenceName = "Go Conference"
	fmt.Println("Welcome to", conferenceName, "booking application") //Assume espacos entre concatenacoes automaticamente
}
