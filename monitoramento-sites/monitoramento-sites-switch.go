package main

import "fmt"

func main(){
    versao := 1
    fmt.Println("Programa para monitoramento de sites")
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("1 - Iniciar Monitoramento")
    fmt.Println("2 - Exibir Logs")
    fmt.Println("0 - Sair do Programa")

    var comando int
    fmt.Scan(&comando)


    fmt.Println("O comando escolhido foi", &comando)
    fmt.Println("O comando escolhido foi", comando)

  switch comando {
  case 1:
       fmt.Println("Monitorando...")
  case 2:
       fmt.Println("Exibindo Logs...")
  case 0:
       fmt.Println("Saindo do programa...")
  default:
        fmt.Println("Não conheço este comando")
    }
 }
