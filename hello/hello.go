package main

import "fmt"
import  "reflect"  

func main(){
  var nome string = "Tamiris"
  var versao float32 = 1.1
  var idade = 37
  cidade := "São Paulo"

  fmt.Println("Olá Mundo,!")
  fmt.Println("Meu Primeiro programa em GO!")
  fmt.Println("Olá Sr(a)." , nome , "e sua idade é", idade)
  fmt.Println("A sua cidade é", cidade)

  fmt.Println("Este programa esta na versão", versao)


  fmt.Println("O tipo da variável idade é", reflect.TypeOf(idade))

}
