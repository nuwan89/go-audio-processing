package main

import (
	"flag"
	"fmt"
	"io"
	"math"

	// "runtime"
	// "sync"
	"time"

	"github.com/hajimehoshi/oto/v2"
)

var (
	sampleRate      = flag.Int("samplerate", 44100, "sample rate")
	channelCount    = flag.Int("channelnum", 2, "number of channel")
	bitDepthInBytes = flag.Int("bitdepthinbytes", 2, "bit depth in bytes")
)

type SineWave struct {
	freq   float64
	length int64
	pos    int64

	remaining []byte
}

func NewSineWave(freq float64, duration time.Duration) *SineWave {
	l := int64(*channelCount) * int64(*bitDepthInBytes) * int64(*sampleRate) * int64(duration) / int64(time.Second)
	l = l / 4 * 4
	return &SineWave{
		freq:   freq,
		length: l,
	}
}

func (s *SineWave) Read(buf []byte) (int, error) {
	if len(s.remaining) > 0 {
		n := copy(buf, s.remaining)
		copy(s.remaining, s.remaining[n:])
		s.remaining = s.remaining[:len(s.remaining)-n]
		return n, nil
	}

	if s.pos == s.length {
		return 0, io.EOF
	}

	eof := false
	if s.pos+int64(len(buf)) > s.length {
		buf = buf[:s.length-s.pos]
		eof = true
	}

	var origBuf []byte
	if len(buf)%4 > 0 {
		origBuf = buf
		buf = make([]byte, len(origBuf)+4-len(origBuf)%4)
	}

	length := float64(*sampleRate) / float64(s.freq)

	num := (*bitDepthInBytes) * (*channelCount)
	p := s.pos / int64(num)
	switch *bitDepthInBytes {
	case 1:
		for i := 0; i < len(buf)/num; i++ {
			const max = 127
			b := int(math.Sin(2*math.Pi*float64(p)/length) * 0.3 * max)
			for ch := 0; ch < *channelCount; ch++ {
				buf[num*i+ch] = byte(b + 128)
			}
			p++
		}
	case 2:
		for i := 0; i < len(buf)/num; i++ {
			const max = 32767
			b := int16(math.Sin(2*math.Pi*float64(p)/length) * 0.3 * max)
			for ch := 0; ch < *channelCount; ch++ {
				buf[num*i+2*ch] = byte(b)
				buf[num*i+1+2*ch] = byte(b >> 8)
			}
			p++
		}
	}

	s.pos += int64(len(buf))

	n := len(buf)
	if origBuf != nil {
		n = copy(origBuf, buf)
		s.remaining = buf[n:]
	}

	if eof {
		return n, io.EOF
	}
	return n, nil
}

func play(context *oto.Context, freq float64, duration time.Duration) oto.Player {
	p := context.NewPlayer(NewSineWave(freq, duration))

	// p := context.NewPlayer(NewSineWave(freq, 1*time.Second))
	p.Play()
	return p
}

func run() error {
	const (
		freqC1 = 440
		freqC  = 523.3
		freqE  = 659.3
		freqG  = 784.0
	)

	c, ready, err := oto.NewContext(*sampleRate, *channelCount, *bitDepthInBytes)
	if err != nil {
		return err
	}

	<-ready

	func() {

		var startTime = time.Now()
		p := play(c, freqC1, 100*time.Hour)

		for true {
			time.Sleep(1 * time.Millisecond)

			if time.Now().Sub(startTime) > 5*time.Second {
				p.Close()
				break
			}

			// fmt.Println("Play")
			// p.Reset()
			// p.Play()
		}
		// time.Sleep(30 * time.Second)
	}()

	fmt.Println("Play callled ended")
	// Pin the players not to GC the players.
	// runtime.KeepAlive(players)
	return nil
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		panic(err)
	}
}
