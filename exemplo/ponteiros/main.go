// ponteiros
package main

import "fmt"

type Carro struct {
	Marca string
}

func (c Carro) MudaMarca(marca string) {
	c.Marca = marca // faz uma cópia da variável e atribui o valor ->>> aponta para outro lugar na memória
	fmt.Println("Mudança 1 " + c.Marca)
}

func (c *Carro) MudaMarca2(marca string) { // aponta diretamente para o local da variável do tipo Carro
	c.Marca = marca // muda o valor de Marca da variável do tipo Carro apontada
	fmt.Println("Mudança 2 " + c.Marca)
}

func main() {
	// Separe um espaço na memória (endereço) e guarde o valor inteiro 1 nele, 
	// a variável a aponta para este endereço
	a := 1 // go infere que valor de a é inteiro e igual a 1
	fmt.Println(a) // imprime o valor de a
	fmt.Println(&a) // imprime o endereço de memória que a aponta
	
	carro := Carro{Marca: "Fiat"}
	carro.MudaMarca("Ford")
	fmt.Println("Após Mudança 1 " + carro.Marca)

	carro.MudaMarca2("Ford")
	fmt.Println("Após Mudança 2 " + carro.Marca)
}
