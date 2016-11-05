package main

import (
    "fmt"
    "io/ioutil"
    "rendergo/tga"
    "rendergo/gfx"
    "rendergo/obj"
)

func main() {

	width := 640
	height := 480

	buffer := gfx.NewBuffer(width,height)
	target := gfx.NewRenderTarget(buffer)

	target.Fill(gfx.Colour{0x00, 0x00, 0x00, 0xFF})

	object := obj.LoadObjFronFile("dat/african_head.obj")

	target.DrawObject(object)

	tgaBytes := tga.CreateTga(width, height, buffer.Pixels() )
	err := ioutil.WriteFile("img.tga", tgaBytes, 0644)
	if(err!=nil) {
		panic(err)
	}

    fmt.Printf("Done\n")
}
