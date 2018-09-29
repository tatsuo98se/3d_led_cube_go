package servicegateway

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"

	zmq "github.com/zeromq/goczmq"
)

var funcIDPlay = "play"
var funcIDPause = "pause"
var funcIDVolume = "volume"
var funcIDStop = "stop"
var funcIDResume = "resume"
var funcIDServiceGatewayAbort = "abort"

type SoundServiceGateway interface {
	Play(soundID string, loop bool, stop bool)
	Pause()
	Stop()
	SetVolume(volume int)
	Terminate()
}

type SoundServiceGatewayImpl struct {
	url       string
	contentID string
	order     chan *soundOrder
	done      chan struct{}
}

type soundOrderData struct {
	ContentID string `json:"content_id"`
	Wav       string `json:"wav"`
	Loop      bool   `json:"loop"`
	AndStop   bool   `json:"and_stop"`
	Val       int    `json:"val"`
}

type soundOrder struct {
	Function string          `json:"func"`
	Data     *soundOrderData `json:"data"`
}

func (s *SoundServiceGatewayImpl) newSoundPlayOrder(soundID string, loop bool, stop bool) *soundOrder {

	o := &soundOrder{}
	o.Function = funcIDPlay

	o.Data = &soundOrderData{}

	o.Data.ContentID = s.contentID
	o.Data.Wav = soundID
	o.Data.Loop = loop
	o.Data.AndStop = stop

	return o
}
func (s *SoundServiceGatewayImpl) newSoundPauseOrder() *soundOrder {
	o := &soundOrder{}
	o.Function = funcIDPause

	o.Data = &soundOrderData{}
	o.Data.ContentID = s.contentID

	return o
}
func (s *SoundServiceGatewayImpl) newSoundResumeOrder() *soundOrder {
	o := &soundOrder{}
	o.Function = funcIDResume

	o.Data = &soundOrderData{}
	o.Data.ContentID = s.contentID

	return o
}
func (s *SoundServiceGatewayImpl) newSoundStopOrder() *soundOrder {
	o := &soundOrder{}
	o.Function = funcIDStop

	o.Data = &soundOrderData{}
	o.Data.ContentID = s.contentID

	return o
}
func (s *SoundServiceGatewayImpl) newSoundVolumeOrder(volume int) *soundOrder {
	o := &soundOrder{}
	o.Function = funcIDVolume

	o.Data = &soundOrderData{}
	o.Data.ContentID = s.contentID
	o.Data.Val = volume

	return o
}
func (s *SoundServiceGatewayImpl) newSoundServiceGatewayAbortOrder() *soundOrder {
	o := &soundOrder{}
	o.Function = funcIDServiceGatewayAbort

	o.Data = &soundOrderData{}
	o.Data.ContentID = s.contentID

	return o
}

var instance SoundServiceGateway
var once sync.Once

func worker(url string, c chan *soundOrder, done chan struct{}) {
	sender, e := zmq.NewPush(url)
	if e != nil {
		log.Fatal(e)
		return
	}
	defer sender.Destroy()
	for {
		order := <-c
		switch order.Function {
		case funcIDPlay:
			fallthrough
		case funcIDPause:
			fallthrough
		case funcIDResume:
			fallthrough
		case funcIDStop:
			fallthrough
		case funcIDVolume:
			if data, e := json.Marshal(order); e == nil {
				fmt.Println(string(data))
				if e := sender.SendFrame(data, zmq.FlagNone); e != nil {
					log.Fatal(e)
				}
			}
		case funcIDServiceGatewayAbort:
			fallthrough
		default:
			// error
			done <- struct{}{}
		}

	}
}

func InitSoundSeriveGateway(url string) {
	rand.Seed(time.Now().UnixNano())
	impl := &SoundServiceGatewayImpl{}

	impl.url = url
	impl.contentID = strconv.Itoa(rand.Int())
	impl.order = make(chan *soundOrder)
	impl.done = make(chan struct{})

	instance = impl
	go worker(impl.url, impl.order, impl.done)
}

func GetSoundSeriveGateway() SoundServiceGateway {
	return instance
}

func (s *SoundServiceGatewayImpl) Play(
	soundID string, loop bool, stop bool) {
	s.order <- s.newSoundPlayOrder(soundID, loop, stop)
}

func (s *SoundServiceGatewayImpl) Pause() {
	s.order <- s.newSoundPauseOrder()
}
func (s *SoundServiceGatewayImpl) Stop() {
	s.order <- s.newSoundStopOrder()
}

func (s *SoundServiceGatewayImpl) SetVolume(volume int) {
	s.order <- s.newSoundVolumeOrder(volume)
}

func (s *SoundServiceGatewayImpl) Terminate() {
	s.order <- s.newSoundServiceGatewayAbortOrder()
	<-s.done
}
