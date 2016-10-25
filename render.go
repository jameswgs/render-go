package main

import (
    "fmt"
    "io/ioutil"
    "rendergo/tga"
    "rendergo/gfx"
    "rendergo/vector"
    "rendergo/obj"
)

func main() {

	width := 640
	height := 480

	buffer := gfx.NewBuffer(width,height)
	target := gfx.NewRenderTarget(buffer)

	target.Fill(gfx.Colour{0x00, 0x00, 0x00, 0xFF})

	target.DrawTri(gfx.NewTri(
			   vector.V2{0,0}, 
			   vector.V2{10,10}, 
			   vector.V2{20,10}, 
			   gfx.Colour{0xFF, 0x00, 0xFF, 0xFF}))

	tgaBytes := tga.CreateTga(width, height, buffer.Pixels() )

	object := obj.LoadObj("https://raw.githubusercontent.com/ssloy/tinyrenderer/master/obj/african_head/african_head.obj")

	target.DrawObject(object)

	err := ioutil.WriteFile("img.tga", tgaBytes, 0644)

	if(err!=nil) {
		panic(err)
	}

    fmt.Printf("Done\n")
}
