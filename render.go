package main

import (
    "fmt"
    "io/ioutil"
    "rendergo/tga"
)

func main() {

	pixels := []byte { 0x00, 0x00, 0xFF, 0xFF } // single white pixel BGRA

	tgaBytes := tga.CreateTga(1, 1, pixels)

	err := ioutil.WriteFile("img.tga", tgaBytes, 0644)

	if(err!=nil) {
		panic(err)
	}

    fmt.Printf("Done\n")
}
