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
	"time"
)

func getUnixNano() int64 {
	return time.Now().UnixNano()
	//    return int64(time.Millisecond)
}

func main() {
	//	runtime.LockOSThread()
	lastUpdate := getUnixNano()
	led := ledlib.NewLedCanvas()
	filter := ledlib.NewLedRollingFilter(led)
	obj := ledlib.NewRocketBitmapObj(filter)
	C.SetUrl(C.CString("192.168.1.12"))
	C.EnableSimulator(false)
	for {
		filter.PreDraw()
		obj.Draw()
		filter.Show()
		current := getUnixNano()
		//		test := ledlib.NewLedCanvas()
		fmt.Println((current - lastUpdate) / (1000 * 1000))
		lastUpdate = current
	}
}
