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
	"net/http"
	_ "net/http/pprof"
)

func getUnixNano() int64 {
	return time.Now().UnixNano()
	//    return int64(time.Millisecond)
}

func main() {
        go func(){
		fmt.Println(http.ListenAndServe("localhost:6060",nil))
	}()
	lastUpdate := getUnixNano()
	led := ledlib.NewLedCanvas()
	obj := ledlib.NewRocketBitmapObj(led)
	C.SetUrl(C.CString("192.168.1.12"))
	C.EnableSimulator(false)
	for {
		obj.Draw()
		led.Show()
		current := getUnixNano()
		//		test := ledlib.NewLedCanvas()
		fmt.Println((current - lastUpdate) / (1000 * 1000))
		lastUpdate = current
	}
}
