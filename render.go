package main

import (
    "fmt"
    "io/ioutil"
    "rendergo/tga"
    "rendergo/gfx"
    "rendergo/obj"
)

func main() {

	object := obj.LoadObjFronFile("dat/african_head.obj")


	width := 640
	height := 480

	buffer := gfx.NewColourBuffer(width,height)
	zbuffer := gfx.NewZBuffer(width,height)
	target := gfx.NewRenderTarget(buffer,zbuffer)

	target.Fill(gfx.Colour{0x00, 0x00, 0x00, 0xFF})
	target.ClearZ()
	target.DrawObject(object)



	tgaBytes := tga.CreateTga(width, height, buffer.Pixels() )
	err := ioutil.WriteFile("img.tga", tgaBytes, 0644)
	if(err!=nil) {
		panic(err)
	}

    fmt.Printf("Done\n")
}
