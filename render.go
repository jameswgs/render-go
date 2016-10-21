package main

import (
    "fmt"
    "io/ioutil"
    "rendergo/tga"
    "rendergo/gfx"
)

func main() {

	width := 2
	height := 2

	buffer := gfx.NewBuffer(width,height)

	buffer.Write( 0, 0, gfx.Colour{ 0xFF, 0x00, 0x00, 0xFF } )
	buffer.Write( 1, 0, gfx.Colour{ 0x00, 0xFF, 0x00, 0xFF } )
	buffer.Write( 0, 1, gfx.Colour{ 0x00, 0x00, 0xFF, 0xFF } )
	buffer.Write( 1, 1, gfx.Colour{ 0x00, 0x00, 0x00, 0xFF } )

	tgaBytes := tga.CreateTga(width, height, buffer.Pixels() )

	err := ioutil.WriteFile("img.tga", tgaBytes, 0644)

	if(err!=nil) {
		panic(err)
	}

    fmt.Printf("Done\n")
}
