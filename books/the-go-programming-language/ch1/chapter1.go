package books

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

// Excercise4 (dup3) prints the count and text of lines that
// appear more than once in the named input files.

// Modify this function to print the names of all files in which each duplicated line occurs.
func Excercise4() {
	fileNameToPrint := ""
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		fileNameToPrint = filename
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n%s", n, line, fileNameToPrint)
		}
	}
}

var palette = []color.Color{color.Black, color.White}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

// Excersise5 can be run by executing `go build cmd/main.go` and then running the resulting binary
// like so: `./main >out.gif`

func Excersise5() {
	// modified because the old random generator was deprecated
	n := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w, n)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	lissajous(os.Stdout, n)
}

func lissajous(out io.Writer, r *rand.Rand) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := r.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
