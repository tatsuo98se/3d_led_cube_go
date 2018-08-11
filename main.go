package main

/*
#cgo LDFLAGS: -lledLib
#include "./lib/led.h"
*/
import "C"
import (
)

func main() {
  for {
    C.SetLed(0,0,0,0xffffff)
    C.Show()
  }
}