package main

/*
#cgo LDFLAGS: -lledLib
#include "./lib/led.h"
*/
import "C"
import (
  "fmt"
  "github.com/tatsuo98se/3d_led_cube_go/ledlib"
)

func main() {
  for {
    C.SetLed(0,0,0,0xffffff)
    C.Show()
    test := new(ledlib.LedCanvas)
    fmt.Println(test)
  }
}