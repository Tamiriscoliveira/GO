package main

import "fmt"
import "os"

func main(){

    exibeIntroducao()
    comando := leComando()

switch comando {
  case 1:
       fmt.Println("Monitorando...")
  case 2:
       fmt.Println("Exibindo Logs...")
  case 0:
       fmt.Println("Saindo do programa...")
       os.Exit(0)
  default:
        fmt.Println("Não conheço este comando")
	os.Exit(-1)
    }
 }


func exibeIntroducao(){
    versao := 1
    fmt.Println("Programa para monitoramento de sites")
    fmt.Println("Este programa está na versão", versao)
    fmt.Println("1 - Iniciar Monitoramento")
    fmt.Println("2 - Exibir Logs")
    fmt.Println("0 - Sair do Programa")
 }


func leComando() int{
    var comandoLido int
    fmt.Scan(&comandoLido)
    fmt.Println("O comando escolhido foi", &comandoLido)
    fmt.Println("O comando escolhido foi", comandoLido)

    return comandoLido
}
