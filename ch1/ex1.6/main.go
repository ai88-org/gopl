package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

// 调色板 由黑色变成绿色
// var palette = []color.Color{color.White, color.Black}
var palette = []color.Color{color.RGBA{R: 0xFF, G: 0, B: 0, A: 0xFF}, color.RGBA{R: 0, G: 0xFF, B: 0, A: 0xFF}, color.RGBA{R: 0, G: 0, B: 0xFF, A: 0xFF},
	color.RGBA{R: 0x33, G: 0x44, B: 0x66, A: 0xFF}}

const (
	whiteIndex = 0 // 没用到
	blackIndex = 2
)

func main() {
	lissajous(os.Stdout)
}
func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 3)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
