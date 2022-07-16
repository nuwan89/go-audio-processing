package main

import (
	"encoding/binary"
	"log"
	"math"
	"time"

	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

const (
	Duration   = 1
	SampleRate = 44100
	Frequency  = 440
)

var (
	tau = math.Pi * 2
)

func _main() {

	var (
		start float64 = 1.0
		end   float64 = 1.0e-4
	)
	nsamps := Duration * SampleRate
	var angle float64 = tau / float64(nsamps)

	decayfac := math.Pow(end/start, 1.0/float64(nsamps))

	var out []byte

	for i := 0; i < nsamps; i++ {
		sample := math.Sin(angle * Frequency * float64(i))
		sample *= start
		start *= decayfac
		var buf [8]byte
		binary.LittleEndian.PutUint32(buf[:], math.Float32bits(float32(sample)))
		append(out[:], buf[:])
	}

	streamer, format, err := wav.Decode(out)
	if err != nil {
		log.Fatal(err)
	}

	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

}
