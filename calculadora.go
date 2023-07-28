package main

import (
    "fmt"
    "strconv"
)

func main() {
    var num1, num2 float64
    fmt.Print("Digite o primeiro número: ")
    fmt.Scanf("%f", &num1)
    fmt.Print("Digite o segundo número: ")
    fmt.Scanf("%f", &num2)

    fmt.Println("Qual operação você deseja realizar?")
    fmt.Println("1. Soma")
    fmt.Println("2. Subtração")
    fmt.Println("3. Multiplicação")
    fmt.Println("4. Divisão")

    var operation int
    fmt.Scanf("%d", &operation)

    switch operation {
    case 1:
        fmt.Println(num1 + num2)
    case 2:
        fmt.Println(num1 - num2)
    case 3:
        fmt.Println(num1 * num2)
    case 4:
        fmt.Println(num1 / num2)
    default:
        fmt.Println("Operação inválida")
    }
}
