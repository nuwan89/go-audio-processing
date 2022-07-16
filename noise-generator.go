package main

import (
	"log"
	"math"
	"math/rand"
	"time"
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

// https://github.com/faiface/beep/wiki/Hello,-Beep!

type Noise struct {
	duration int64
	start time.Time
	pos   int64
	shouldPlay bool
}

var pos = 0
var start time.Time = time.Now()

func (n *Noise) Stream(samples [][2]float64) (x int, ok bool) {

	// var (
	// 	start float64 = 1.0
	// 	end   float64 = 1.0e-4
	// )
	nsamps := n.duration * SampleRate
	// fmt.Println(nsamps)
	// var angle float64 = math.Pi * 2 / float64(nsamps)

	// decayfac := math.Pow(end/start, 1.0/float64(nsamps))
	// for i := n.pos; i < nsamps; i++ {
		
		for i := range samples {
			if n.pos < nsamps && n.shouldPlay {
				if !n.shouldPlay {
					// return 0, false
					return len(samples), false
				}
				samples[i][0] = rand.Float64()*2 - 1
				samples[i][1] = rand.Float64()*2 - 1
				n.pos += 1
		} else { 
			return len(samples), false
		}
	}
	return len(samples), true

	// for i := range samples {
	// 	samples[i][0] = rand.Float64()*2 - 1
	// 	samples[i][1] = rand.Float64()*2 - 1
	// 	pos += 1
	// 	n.pos += 1
	// }
	// log.Println(start,"duration:", time.Now().Sub(start).Seconds(), "S samples:", pos,"Rate:", float64(pos)/time.Now().Sub(start).Seconds())
	// return len(samples), true
}

const (
	Duration   = 1
	SampleRate = 44100
	// SampleRate = 192000
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
	log.Println("Start:", time.Now())
	sr := beep.SampleRate(SampleRate)
	i := 0
	for {
		if i > 100 {
			fmt.Println("Stoped")
			break
		}
		// fmt.Println(i)
		i += 1
		// time.Sleep(1 * time.Millisecond)
		time.Sleep(1 * time.Microsecond)
	}
	if 1 == 1 {
		panic("ssssss")
	}
	// N stands for number), which calculates the number of samples contained in the provided duration.
	speaker.Init(sr, sr.N(time.Second/10)) // sr.N(time.Second/10) = buffer size for duration 1/10 second
	done := make(chan bool)
	streamer := Noise{start: time.Now(), duration: 5, shouldPlay: true}
	speaker.Play(beep.Seq(&streamer, beep.Callback(func(){
		done <- true
	})))


	for {
		if i > 100 {
			streamer.shouldPlay = false
			fmt.Println("Stoped")
			break
		}

		// select {
		// // case <-done:
		// // 	return
		// 	case <-time.After(time.Millisecond):

				// speaker.Lock()
				// fmt.Println(sr.D(streamer.Position()).Round(time.Second))
				// fmt.Println(i)
				// speaker.Unlock()
			// }
		// fmt.Println(i)
		i += 1
		time.Sleep(1 * time.Millisecond)
	}
	
	// select {} // makes the program hang forever
}



// func generate(samplerate, freq, duration) {
// 	var (
// 		start float64 = 1.0
// 		end   float64 = 1.0e-4
// 	)
// 	nsamps := duration * samplerate
// 	var angle float64 = math.Pi * 2 / float64(nsamps)

// 	decayfac := math.Pow(end/start, 1.0/float64(nsamps))
// 	for i := 0; i < nsamps; i++ {
// 		sample := math.Sin(angle * freq * float64(i))
// 		sample *= start
// 		start *= decayfac
// 		var buf [8]byte
// 		binary.LittleEndian.PutUint32(buf[:], math.Float32bits(float32(sample)))
// 		// bw, err := f.Write(buf[:])
// 		// if err != nil {
// 		// 	panic(err)
// 		// }
// 		fmt.Printf("\rWrote: %v bytes to %s", bw, file)
// 	}
// }
