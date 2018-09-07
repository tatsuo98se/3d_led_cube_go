package ledlib

/*
#cgo LDFLAGS: -lledlib
#include "./../../lib/led.h"
*/
import "C"
import (
	"log"
	"net"
	"strconv"
)

const LedWidth = 16
const LedHeight = 32
const LedDepth = 8

const ledColor = 3
const ledRed = 0
const ledGreen = 1
const ledBulue = 2

type Led interface {
	SetUrl(url string)
	SetLed(x, y, z int, rgb uint32)
	Clear()
	Show()
	EnableSimulator(enable bool)
	SetPort(port uint16)
}

var sharedInstance Led

func GetLed() Led {
	if sharedInstance == nil {
		sharedInstance = newLed()
	}
	return sharedInstance
}

func newLed() *ledImpl {
	goImpl := newGoLed()
	cImpl := newCLed()
	return &ledImpl{goImpl, cImpl, cImpl, true}
}

/*
* ledImpl
 */
type ledImpl struct {
	goImpl          *ledGoImpl
	cImpl           *ledCImpl
	currentImpl     Led
	enableSimulator bool
}

func (led *ledImpl) SetUrl(url string) {
	led.currentImpl.SetUrl(url)
}

func (led *ledImpl) SetLed(x, y, z int, rgb uint32) {
	led.currentImpl.SetLed(x, y, z, rgb)
}

func (led *ledImpl) Clear() {
	led.currentImpl.Clear()
}

func (led *ledImpl) Show() {
	led.currentImpl.Show()
}

func (led *ledImpl) EnableSimulator(enable bool) {
	if enable {
		led.currentImpl = led.cImpl
	} else {
		led.currentImpl = led.goImpl
	}
	C.EnableSimulator(C.bool(enable))
}

func (led *ledImpl) SetPort(port uint16) {
	led.currentImpl.SetPort(port)
}

/*
* Go Implimentation
 */
type ledGoImpl struct {
	ledUrl       string
	ledPort      uint16
	led565Buffer []byte
	sem          chan struct{}
}

func newGoLed() *ledGoImpl {
	led := ledGoImpl{}
	led.led565Buffer = make([]byte, LedWidth*LedHeight*LedDepth*2)

	led.sem = make(chan struct{}, 1)
	return &led
}

func (led *ledGoImpl) SetUrl(url string) {
	led.ledUrl = url
}

func (led *ledGoImpl) SetLed(x, y, z int, rgb uint32) {
	if x < 0 || LedWidth <= x {
		log.Fatalf("invalid x : %d\n", x)
		return
	}
	if y < 0 || LedHeight <= y {
		log.Fatalf("invalid y : %d\n", y)
		return
	}
	if z < 0 || LedDepth <= z {
		log.Fatalf("invalid z : %d\n", z)
		return
	}

	r, g, b := byte(rgb>>16), byte(rgb>>8), byte(rgb>>0)

	index565 := z*2 + y*LedDepth*2 + x*LedHeight*LedDepth*2
	led.led565Buffer[index565+0] = r&0xF8 + g>>5
	led.led565Buffer[index565+1] = (g<<2)&0xe0 + b>>3
}

func (led *ledGoImpl) Clear() {
	for i, _ := range led.led565Buffer {
		led.led565Buffer[i] = 0
	}
}

func (led *ledGoImpl) Show() {
	tcpAddr, err := net.ResolveUDPAddr("udp", led.getUrl())
	if err != nil {
		log.Fatalf("error: %s", err.Error())
		return
	}
	conn, err := net.DialUDP("udp", nil, tcpAddr)
	if err != nil {
		log.Fatalf("error: %s", err.Error())
		return
	}
	defer conn.Close()
	conn.Write(led.led565Buffer)
}

func (led *ledGoImpl) EnableSimulator(enable bool) {
	// do nothing.
}

func (led *ledGoImpl) SetPort(port uint16) {
	led.ledPort = port
}

func (led *ledGoImpl) getUrl() string {
	if led.ledPort == 0 {
		return led.ledUrl
	} else {
		return led.ledUrl + ":" + strconv.FormatUint(uint64(led.ledPort), 10)
	}
}

type ledCImpl struct {
}

func newCLed() *ledCImpl {
	return &ledCImpl{}
}

func (led *ledCImpl) SetUrl(url string) {
	C.SetUrl(C.CString(url))
}

func (led *ledCImpl) SetLed(x, y, z int, rgb uint32) {
	C.SetLed(C.int(x), C.int(y), C.int(z), C.int(rgb))
}

func (led *ledCImpl) Clear() {
	C.Clear()
}

func (led *ledCImpl) Show() {
	C.Show()
}

func (led *ledCImpl) EnableSimulator(enable bool) {
	C.EnableSimulator(C.bool(enable))
}

func (led *ledCImpl) SetPort(port uint16) {
	C.SetPort(C.ushort(port))
}
