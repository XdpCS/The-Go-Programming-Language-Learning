package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

type lissajousConfig struct {
	cycles  float64
	res     float64
	size    int
	nframes int
	delay   int
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajousConfig := newLissajousConfig()
		queries := r.URL.Query()
		for queryK, queryV := range queries {
			var err error
			switch queryK {
			case "cycles":
				lissajousConfig.cycles, err = strconv.ParseFloat(queryV[0], 64)
			case "res":
				lissajousConfig.res, err = strconv.ParseFloat(queryV[0], 64)
			case "size":
				lissajousConfig.size, err = strconv.Atoi(queryV[0])
			case "nframes":
				lissajousConfig.nframes, err = strconv.Atoi(queryV[0])
			case "delay":
				lissajousConfig.delay, err = strconv.Atoi(queryV[0])
			}

			if err != nil {
				log.Fatal("Invalid query parameter: " + queryK + " = " + queryV[0] + "error: " + err.Error() + "\n")
			}
		}
		lissajous(w, lissajousConfig)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, lissajousConfig *lissajousConfig) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: lissajousConfig.nframes}
	phase := 0.0 // phase difference
	for i := 0; i < lissajousConfig.nframes; i++ {
		rect := image.Rect(0, 0, 2*lissajousConfig.size+1, 2*lissajousConfig.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < lissajousConfig.cycles*2*math.Pi; t += lissajousConfig.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(lissajousConfig.size+int(x*float64(lissajousConfig.size)+0.5), lissajousConfig.size+int(y*float64(lissajousConfig.size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, lissajousConfig.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func newLissajousConfig() *lissajousConfig {
	return &lissajousConfig{
		cycles:  5,
		res:     0.001,
		size:    100,
		nframes: 64,
		delay:   8,
	}
}
