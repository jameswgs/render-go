package main

import (
    "fmt"
    "io/ioutil"
    "rendergo/tga"
    "rendergo/gfx"
    "rendergo/vector"
)

func main() {

	width := 20
	height := 20

	buffer := gfx.NewBuffer(width,height)

	tri := gfx.NewTri(
			   vector.V2{10,0}, 
			   vector.V2{0,10}, 
			   vector.V2{20,10}, 
			   gfx.Colour{0xFF, 0x00, 0xFF, 0xFF})

	buffer.Write( 0, 0, gfx.Colour{ 0xFF, 0x00, 0x00, 0xFF } )
	buffer.Write( 1, 0, gfx.Colour{ 0x00, 0xFF, 0x00, 0xFF } )
	buffer.Write( 0, 1, gfx.Colour{ 0x00, 0x00, 0xFF, 0xFF } )
	buffer.Write( 1, 1, gfx.Colour{ 0x00, 0x00, 0x00, 0xFF } )

	target := gfx.NewRenderTarget(buffer)

	target.Fill(gfx.Colour{0x00, 0x00, 0x00, 0xFF})

	target.Draw(tri)


	tgaBytes := tga.CreateTga(width, height, buffer.Pixels() )

	err := ioutil.WriteFile("img.tga", tgaBytes, 0644)

	if(err!=nil) {
		panic(err)
	}

    fmt.Printf("Done\n")
}
