package main

/*
#cgo LDFLAGS: -lledLib
#include "./lib/led.h"
*/
import "C"
import (
  "fmt"
//  "github.com/tatsuo98se/3d_led_cube_go/ledlib"
  "time"
)

func getUnixNano() int64{
    return time.Now().UnixNano()
//    return int64(time.Millisecond)
}

func main() {
  lastUpdate := getUnixNano()
  C.EnableSimulator(false)
  for {
    C.SetLed(0,0,0,0xffffff)
    C.Show()
    current := getUnixNano()
//    test := ledlib.NewLedCanvas()
    fmt.Println( (current - lastUpdate)/(1000 * 1000) )
    lastUpdate = current
  }
}