package main

import (
    "fmt"
    "io/ioutil"
    "rendergo/tga"
    "rendergo/gfx"
//    "rendergo/vector"
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

	// target.DrawTri(
	// 	gfx.NewTri(
	// 		//vector.V2{275, 200}, vector.V2{297, 188}, vector.V2{280, 213}, 
	// 		vector.V2{363, 205}, vector.V2{362, 209}, vector.V2{367, 208},
	// 		gfx.Colour{0xFF, 0xFF, 0xFF, 0xFF}))

	tgaBytes := tga.CreateTga(width, height, buffer.Pixels() )
	err := ioutil.WriteFile("img.tga", tgaBytes, 0644)
	if(err!=nil) {
		panic(err)
	}

    fmt.Printf("Done\n")
}
