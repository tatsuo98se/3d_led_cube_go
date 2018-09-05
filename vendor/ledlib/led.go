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

const LED_WIDTH = 16
const LED_HEIGHT = 32
const LED_DEPTH = 8
const LED_COLOR = 3
const LED_RED = 0
const LED_GREEN = 1
const LED_BLUE = 2

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
	ledUrl    string
	ledPort   uint16
	ledBuffer []byte
	sem       chan struct{}
}

func newGoLed() *ledGoImpl {
	led := ledGoImpl{}
	led.ledBuffer = make([]byte, LED_WIDTH*LED_HEIGHT*LED_DEPTH*LED_COLOR)
	led.sem = make(chan struct{}, 1)
	return &led
}

func (led *ledGoImpl) SetUrl(url string) {
	led.ledUrl = url
}

func (led *ledGoImpl) SetLed(x, y, z int, rgb uint32) {
	if x < 0 || LED_WIDTH <= x {
		log.Fatalf("invalid x : %d\n", x)
		return
	}
	if y < 0 || LED_HEIGHT <= y {
		log.Fatalf("invalid y : %d\n", y)
		return
	}
	if z < 0 || LED_DEPTH <= z {
		log.Fatalf("invalid z : %d\n", z)
		return
	}
	index := z*LED_COLOR + y*LED_DEPTH*LED_COLOR + x*LED_HEIGHT*LED_DEPTH*LED_COLOR
	led.ledBuffer[index+LED_RED] = byte(rgb >> 16)
	led.ledBuffer[index+LED_GREEN] = byte(rgb >> 8)
	led.ledBuffer[index+LED_BLUE] = byte(rgb >> 0)
}

func (led *ledGoImpl) Clear() {
	for i, _ := range led.ledBuffer {
		led.ledBuffer[i] = 0
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
	udpBuffer := rgb888toRGB565(led.ledBuffer)
	conn.Write(udpBuffer)
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

func (led *ledGoImpl) getLedBuffer() []byte {
	return led.ledBuffer
}

func rgb888toRGB565(rgb888 []byte) []byte {
	const RGB888_RGB_SIZE = 3
	const RGB565_RGB_SIZE = 2

	lengthOfRGB888 := len(rgb888)
	lengthOfRGB565 := lengthOfRGB888 / RGB888_RGB_SIZE * RGB565_RGB_SIZE
	rgb565 := make([]byte, lengthOfRGB565)
	for i := 0; i < len(rgb888); i += RGB888_RGB_SIZE {
		r := rgb888[i]
		g := rgb888[i+1]
		b := rgb888[i+2]
		indexOfRGB565 := i / RGB888_RGB_SIZE * RGB565_RGB_SIZE
		rgb565[indexOfRGB565] = r&0xF8 + g>>5
		rgb565[indexOfRGB565+1] = g&0x1C + b>>3
	}
	return rgb565
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
