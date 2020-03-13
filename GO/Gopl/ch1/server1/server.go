package main

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
	"strconv"
)

var palette = []color.Color{color.White,color.RGBA{0, 255, 68, 255}, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

//var mu sync.Mutex
//var count int

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//处理http请求
func handler(w http.ResponseWriter, r *http.Request) {
	//if语句中嵌套表达式
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		if k == "cycles" {
			//函数有多个返回值时，multiple-value strconv.Atoi() in single-value context ，不需要的参数要用_接收
			cycles, _ := strconv.ParseFloat(v[0], 64)
			//fmt.Fprintf(w,"%q:%d",k,cycles);
			//调用gif图形函数
			if cycles < 50 {
				lissajousA(w, cycles)
			}
		} else {
			_, _ = fmt.Fprintf(w, "URL PATH:%q <br/>", r.URL.Path)
			_, _ = fmt.Fprintf(w, "%q:%q", k, v)
		}
	}
}

func lissajousA(out io.Writer, cycles float64)  {
	const (
		res		= 0.01
		size	= 100
		nFrames	= 64
		delay	= 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount:nFrames}

	phase := 0.0 // phase difference
	for i := 0; i < nFrames; i++ {
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
	_ = gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}


