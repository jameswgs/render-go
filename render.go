package main

import (
    "fmt"
    "io/ioutil"
)

func createTga(w uint16, h uint16, pixels []byte) []byte {

	// TGA FORMAT
	// no ID size: 0x00 
	// no colour map: 0x00
	// image type, uncompressed true colour: 0x02
	// colour map spec (none): 
		// first index: short: 0x0000
		// length: short: 0x0000
		// bbp: 0x00
	// image spec:
		// x-origin: short: 0x0000
		// y-origin: short: 0x0000
		// width: short: 0x0001
		// height: short: 0x0001
		// bpp: 0x04
		// alpha type & depth - 2 nibbles: 0x38
	// colour information:
		// 1x1 32bit white pixel: 0xFFFFFFFF

	if(len(pixels)!=(int)(4*w*h)) {
		panic("pixels need to be of size 4 * width * height")
	}

	w0 := (byte) (w&0xFF)
	w1 := (byte) (w>>8)

	h0 := (byte) (h&0xFF)
	h1 := (byte) (h>>8)

	tgaHeader := []byte {
		0x00, // id length
		0x00, // colour map size
		0x02, // uncompressed true colour
		0x00, 0x00, 0x00, 0x00, 0x00, // colour map spec (none)
		0x00, 0x00, // x origin 
		0x00, 0x00, // y-origin
		w0, w1, // width little endian 1px
		h0, h1, // height little endian 1px
		0x20, // 32 bpp
		0x38, // alpha type and alpha depth
		} // single white pixel

	return append(tgaHeader, pixels...)

}

func main() {

	pixels := []byte { 0x00, 0x00, 0xFF, 0xFF } // single white pixel BGRA

	tgaBytes := createTga(1, 1, pixels)

	err := ioutil.WriteFile("img.tga", tgaBytes, 0644)

	if(err!=nil) {
		panic(err)
	}

    fmt.Printf("Done\n")
}
