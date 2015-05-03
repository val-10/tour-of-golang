package main

import ("golang.org/x/tour/pic"
		"image/color"
		"image")

type Image struct{w, h int}

func (m Image) ColorModel() color.Model{
	return color.RGBAModel
}
func (m Image) Bounds() image.Rectangle{
	return image.Rect(0, 0, m.w, m.h)
}
func (m Image) At(x, y int) color.Color{
	v:=uint8(x*x+y*y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{255,255}
	pic.ShowImage(m)
}