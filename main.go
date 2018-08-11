package main

/*
#cgo LDFLAGS: -lledLib
#include "./lib/led.h"
*/
import "C"
import (
  "github.com/3d_led_cube_go/ledlib/led_canvas"
)

func main() {
  for {
    C.SetLed(0,0,0,0xffffff)
    C.Show()
  }
}