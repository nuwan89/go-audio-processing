package main

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

// https://github.com/faiface/beep/wiki/Hello,-Beep!

type Noise struct {
	start time.Time
	pos   int64
}

func (n Noise) Stream(samples [][2]float64) (x int, ok bool) {

	log.Println("duration:", time.Now().Sub(n.start)/time.Second, "S samples:", n.pos)
	for i := range samples {
		samples[i][0] = rand.Float64()*2 - 1
		samples[i][1] = rand.Float64()*2 - 1
		n.pos += 1
	}
	return len(samples), true
}

const (
	Duration   = 1
	SampleRate = 44100
	Frequency  = 440
)

var (
	tau = math.Pi * 2
)

// func sine() float64 {

// }

func (n Noise) Err() error {
	return nil
}

func main() {

	sr := beep.SampleRate(SampleRate)
	// N stands for number), which calculates the number of samples contained in the provided duration.
	speaker.Init(sr, sr.N(time.Second/10)) // sr.N(time.Second/10) = buffer size for duration 1/10 second
	speaker.Play(Noise{start: time.Now()})
	select {} // makes the program hang forever
}
