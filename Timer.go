package main

import (
  "fmt"
  "math/rand"
  "time"
  //"os"
  //"bufio"
  "github.com/MarinX/keylogger"
  "github.com/sirupsen/logrus"
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

func UpdateClock(s, ms int) (int, int){
  fmt.Print("\033[u")
  if ms < 10 {
    fmt.Print("\033[K\r", s,":0",ms)
  } else if ms < 100 {
    fmt.Print("\033[K\r", s,":",ms)
  } else {
    ms = 0
    s++
    fmt.Print("\033[K\r", s,":0",ms)
  }
  fmt.Print("\033[2B")
  ms++
  time.Sleep(8871 * time.Microsecond)
  return s, ms
}


func Stopwatch(k *keylogger.KeyLogger) (int, int){
  var ms int
  var s int

  fmt.Print("\n\033[s") // save timer position
    for {
      event := k.Read()
      for e := range event {
        if e.Type == keylogger.EvKey {
          if e.KeyString() != "SPACE" {
            fmt.Print("\033[K\r")
            s, ms = UpdateClock(s, ms)
          } else if e.KeyString() == "SPACE" {
            return s, ms
          }
        }
      }
    }
    return s, ms
}


func main() {
  rand.Seed(time.Now().UnixNano())
  keyboard := keylogger.FindKeyboardDevice()

  // check if we found a path to keyboard
  if len(keyboard) <= 0 {
    logrus.Error("No keyboard found...you will need to provide manual input path")
    return
  }

  logrus.Println("Found a keyboard at", keyboard)
  // init keylogger with keyboard
  k, err := keylogger.New(keyboard)
  if err != nil {
    logrus.Error(err)
    return
  }
  defer k.Close()
  Scramble()
  //input := k.Read().Type
  event := k.Read()
  for {
    for e := range event {
      if e.Type == keylogger.EvKey {
        if e.KeyString() == "SPACE" {
          for i := 0; i < 1000; i++ {
            if e.KeyString() != "SPACE" {
              break
              break
            }
            time.Sleep(1 * time.Millisecond)
          }
          event = k.Read()
          for i := range event {
            if i.Type == keylogger.EvKey {
              if i.KeyRelease() == true {
                s, ms := Stopwatch(k)
                fmt.Println("\n",s,":",ms)
              }
            }
          }
        }
      }
    }
  }

}

/*
func main() {
  // find keyboard device, does not require a root permission
	keyboard := keylogger.FindKeyboardDevice()

	// check if we found a path to keyboard
	if len(keyboard) <= 0 {
		logrus.Error("No keyboard found...you will need to provide manual input path")
		return
	}

	logrus.Println("Found a keyboard at", keyboard)
	// init keylogger with keyboard
	k, err := keylogger.New(keyboard)
	if err != nil {
		logrus.Error(err)
		return
	}

	defer k.Close()
  events := k.Read()

  for e := range events {
    switch e.Type {
      case keylogger.EvKey:
      if e.KeyPress() {
        logrus.Println("[event] press key", e.KeyString())
      }
      if e.KeyPress() {
        logrus.Println("[event] release key", e.KeyString())
      }
      break
    }
  }
}
*/
