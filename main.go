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
	_ "net/http/pprof"
	"time"
)

func getUnixNano() int64 {
	return time.Now().UnixNano()
	//    return int64(time.Millisecond)
}

func main() {
	//	runtime.LockOSThread()
	go func() {
		//		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	lastUpdate := getUnixNano()
	led := ledlib.NewLedCanvas()
	filter1 := ledlib.NewLedRollingFilter(led)
	filter2 := ledlib.NewLedSkewedFilter(filter1)
	obj := ledlib.NewRocketBitmapObj()
	C.SetUrl(C.CString("192.168.1.12"))
	C.EnableSimulator(false)
	for {
		filter2.PreShow()
		obj.Draw(filter2)
		current := getUnixNano()
		//		test := ledlib.NewLedCanvas()
		fmt.Println((current - lastUpdate) / (1000 * 1000))
		lastUpdate = current
	}
}
