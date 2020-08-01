package main

import (
  "fmt"
  "math/rand"
  "time"
  //"os"
  //"bufio"
  "github.com/MarinX/keylogger"
)

func Scramble() {
  Arr := make([]int, 20)
  move := make([]int, 20)
  form := make([]int, 20)
  for i := range Arr {
    for move[i] == 0 || move[i] == move[i-1] { //prevent situations like R R'
      move[i] = rand.Intn(6)
      if i == 0 {
        break
      }
    }
    form[i] = rand.Intn(3)
    Arr[i] = 3 * move[i] + form[i]
  }

  fmt.Print("Scramble: ")

  for i := range Arr {
    switch Arr[i] { // each possible move
      case 0 :
        fmt.Print("R ")
      case 1 :
        fmt.Print("R' ")
      case 2 :
        fmt.Print("R2 ")
      case 3 :
        fmt.Print("L ")
      case 4 :
        fmt.Print("L' ")
      case 5 :
        fmt.Print("L2 ")
      case 6 :
        fmt.Print("U ")
      case 7 :
        fmt.Print("U' ")
      case 8 :
        fmt.Print("U2 ")
      case 9 :
        fmt.Print("D ")
      case 10 :
        fmt.Print("D' ")
      case 11 :
        fmt.Print("D2 ")
      case 12 :
        fmt.Print("F ")
      case 13 :
        fmt.Print("F' ")
      case 14 :
        fmt.Print("F2 ")
      case 15 :
        fmt.Print("B ")
      case 16 :
        fmt.Print("B' ")
      case 17 :
        fmt.Print("B2 ")
    }
  }
  fmt.Print("\n")
}


func Stopwatch() {
  var ms int
  var s int
  fmt.Print("\033[s","\n\n")
  var input string
  fmt.Scanf(input)
  for input != " "{
    fmt.Print("\033[2A")
    if ms < 10 {
      fmt.Print("\033[K\r", s,":0",ms)
    } else if ms < 100 {
      fmt.Print("\033[K\r", s,":",ms)
    } else {
      ms = 0
      s++
      fmt.Print("\033[K\r", s,":0",ms)
    }

    fmt.Print("\033[2B", "\033[u")
    ms++
    //fmt.Print("\033[1A", "\033[1B", "0","\0330")
    time.Sleep(10871 * time.Microsecond)
  }

}


func main() {
  rand.Seed(time.Now().UnixNano())
  Scramble()
  //scan := bufio.NewReader(os.Stdin)
  //go Stopwatch()
  //input, _ := scan.ReadByte()
  var input string
  fmt.Scanf(input)
  for input != " " {
    //input, _ = scan.ReadByte()
    fmt.Scanf(input)
    //fmt.Println()
  }

  fmt.Println("\nDone2")
}
