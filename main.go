

package main

import "fmt"
import "math"

func Sum (firstNumber float64, secondNumber float64){
  operationResult := firstNumber + secondNumber;
  fmt.Println("\nResultado: ", operationResult, "\n"); 
}

func Subtract (firstNumber float64, secondNumber float64){
  operationResult := firstNumber - secondNumber;
  fmt.Println("\nResultado: ", operationResult, "\n"); 
}

func Multiply (firstNumber float64, secondNumber float64){
  operationResult := firstNumber * secondNumber;
  fmt.Println("\nResultado: ", operationResult, "\n"); 
}

func Divide (firstNumber float64, secondNumber float64){
  if (secondNumber != 0) {
    operationResult := firstNumber / secondNumber;
    fmt.Println("\nResultado: ", operationResult, "\n");   
  } else {
    fmt.Println("\nDivisões por zero não são permitidas\n");   
  }
}

func Square (number float64){
  operationResult := number * number ;
  fmt.Println("\nResultado: ", operationResult, "\n"); 
}

func Cube (number float64){
  operationResult := (number * number) * number ;
  fmt.Println("\nResultado: ", operationResult, "\n"); 
}

func Squareroot(number float64){
  operationResult := math.Sqrt(number);
  fmt.Println("\nResultado: ", operationResult, "\n"); 
}

func Cubicroot(number float64){
  operationResult := math.Cbrt(number);
  fmt.Println("\nResultado: ", operationResult, "\n"); 
}


func Log (number float64){
  operationResult := math.Log10(number) ;
  fmt.Println("\nResultado: ", operationResult, "\n"); 
}



func main() {
  var running bool = true;
  var operation int = 0;
  
  
  for running {
    var firstNumber float64 = 0; 
    var secondNumber float64 = 0;
    fmt.Println("1 - Soma\n2 - Subtração\n3 - Multiplicação\n4 - Divisão\n5 - Elevar ao quadrado\n6 - Elevar ao cubo\n7 - Raiz quadrada\n8 - Raiz cubica\n9 - Logaritmo\n10 - Sair")
    fmt.Println("\nInsira o número que representa a operação que deseja fazer: ")
    fmt.Scanf("%d", &operation)
    switch operation {
      case 1:
        fmt.Println("Insira os operandos")
        fmt.Scanf("%f %f", &firstNumber, &secondNumber)
        Sum(firstNumber, secondNumber)
      case 2:
        fmt.Println("Insira os operandos")
        fmt.Scanf("%f %f", &firstNumber, &secondNumber)
        Subtract(firstNumber, secondNumber)
      case 3:
        fmt.Println("Insira os operandos")
        fmt.Scanf("%f %f", &firstNumber, &secondNumber)
        Multiply(firstNumber, secondNumber)
      case 4:
        fmt.Println("Insira os operandos")
        fmt.Scanf("%f %f", &firstNumber, &secondNumber)
        Divide(firstNumber, secondNumber)
      case 5:
        fmt.Println("Insira o operando")
        fmt.Scanf("%f", &firstNumber)
        Square(firstNumber)
      case 6:
        fmt.Println("Insira o operando")
        fmt.Scanf("%f", &firstNumber)
        Cube(firstNumber)
      case 7:
        fmt.Println("Insira o operando")
        fmt.Scanf("%f", &firstNumber)
        Squareroot(firstNumber)
      case 8:
        fmt.Println("Insira o operando")
        fmt.Scanf("%f", &firstNumber)
        Cubicroot(firstNumber)
      case 9:
        fmt.Println("Insira o operando")
        fmt.Scanf("%f", &firstNumber)
        Log(firstNumber)
      case 10:
        fmt.Println("Saindo...")
        running = false;
      default: 
        fmt.Println("\nPor favor, insira um número valido\n")
    } 
  }
}