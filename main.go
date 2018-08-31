package main

/*
#cgo LDFLAGS: -lledlib
#include "./lib/led.h"
*/
import "C"
import (
	"github.com/tatsuo98se/3d_led_cube_go/ledlib"
	//  "fmt"

	"fmt"
	"runtime"
	"time"
)

func getUnixNano() int64 {
	return time.Now().UnixNano()
	//    return int64(time.Millisecond)
}

func main() {
	runtime.LockOSThread()
	lastUpdate := getUnixNano()
	led := ledlib.NewLedCanvas()
	obj := ledlib.NewRocketBitmapObj(led)
	C.EnableSimulator(true)
	for {
		obj.Draw()
		led.Show()
		current := getUnixNano()
		//		test := ledlib.NewLedCanvas()
		fmt.Println((current - lastUpdate) / (1000 * 1000))
		lastUpdate = current
	}
}
