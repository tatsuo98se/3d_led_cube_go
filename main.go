package main

/*
#cgo LDFLAGS: -lledlib
#include "./lib/led.h"
*/
import "C"
import (
	"flag"
	"fmt"
	"ledlib"
	_ "net/http/pprof"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func getUnixNano() int64 {
	return time.Now().UnixNano()
}

func main() {
	var (
		optDestination = flag.String("d", "", "Specify IP and port of Led Cube. if opt is not set, launch simulator.")
	)
	flag.Parse()
	if *optDestination == "" {
		runtime.LockOSThread()
		C.EnableSimulator(true)
	} else {
		C.EnableSimulator(false)
		ipAndPort := strings.Split(*optDestination, ":")
		switch {
		case len(ipAndPort) == 2:
			C.SetUrl(C.CString(ipAndPort[0]))
			port, e := strconv.ParseInt(ipAndPort[1], 10, 16)
			if e != nil {
				fmt.Printf("invalid port number. %s", ipAndPort[1])
				return
			}
			C.SetPort(C.ushort(port))
		case len(ipAndPort) == 1:
			C.SetUrl(C.CString(*optDestination))
		case len(ipAndPort) == 0:
			// do nothing
		default:
			fmt.Println("")
			return
		}

	}
	fmt.Println(*optDestination)
	go func() {
		//		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	lastUpdate := getUnixNano()
	led := ledlib.NewLedCanvas()
	filter1 := ledlib.NewLedRollingFilter(led)
	filter2 := ledlib.NewLedSkewedFilter(filter1)
	obj := ledlib.NewRocketBitmapObj()

	for {
		filter2.PreShow()
		obj.Draw(filter2)
		current := getUnixNano()
		fmt.Println((current - lastUpdate) / (1000 * 1000))
		lastUpdate = current
	}
}
