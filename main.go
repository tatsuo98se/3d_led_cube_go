package main

import (
	"bytes"
	"flag"
	"fmt"
	"ledlib"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
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
		ledlib.GetLed().EnableSimulator(true)
	} else {
		ledlib.GetLed().EnableSimulator(false)
		ledlib.GetLed().SetUrl(*optDestination)
	}
	fmt.Println(*optDestination)
	go func() {
		//		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	/*
		setup renderer
	*/
	renderer := ledlib.NewLedBlockRenderer()
	renderer.Start()

	// start http server
	// endpoins are below
	// POST /api/show       content json
	// POST /api/abort		no content
	// POST /api/target     text content
	//
	http.HandleFunc("/api/show", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			bufbody := new(bytes.Buffer)
			bufbody.ReadFrom(r.Body)
			fmt.Fprintln(w, bufbody.String())
			renderer.Show(bufbody.String())
		default:
			http.Error(w, "Not implemented.", http.StatusNotFound)
		}
	})
	http.HandleFunc("/api/abort", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			fmt.Fprintln(w, "abort")
			renderer.Abort()
		default:
			http.Error(w, "Not implemented.", http.StatusNotFound)
		}
	})
	http.HandleFunc("/api/target", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
		default:
			http.Error(w, "Not implemented.", http.StatusNotFound)
		}
	})
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "Hello")
		default:
			http.Error(w, "Not implemented.", http.StatusNotFound)
		}
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
	renderer.Terminate()
	/*
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
		}*/

	/*
		renderer.Show(`{"orders":[{"id":"filter-rolling"},{"id":"object-rocket", "lifetime":1},{"id":"filter-skewed"},{"id":"object-rocket"}]}`)
		time.Sleep(3 * time.Second)
		renderer.Show(`{"orders":[{"id":"filter-rolling"},{"id":"object-rocket", "lifetime":1},{"id":"filter-skewed"},{"id":"object-rocket"}]}`)
		time.Sleep(10 * time.Second)
		renderer.Abort()
		time.Sleep(3 * time.Second)
		renderer.Terminate()*/
}
