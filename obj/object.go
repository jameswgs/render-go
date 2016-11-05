package obj

import (
	"os"
	"io"
	"bufio"
	"net/http"
	"strings"
	"strconv"
	"rendergo/vector"
)

type Face struct {
	A, B, C int
}

type Object struct {
	Verts []vector.V3f
	Faces []Face
}

func LoadObjFronUrl(url string) *Object {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	return LoadObjFromReader(response.Body)	
}

func LoadObjFronFile(fileName string) *Object {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()
	return LoadObjFromReader(file)	
}

func LoadObjFromReader(reader io.Reader) *Object {

	scanner := bufio.NewScanner(reader)

	verts := make([]vector.V3f, 0)
	faces := make([]Face, 0)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		switch(parts[0]) {
		case "v" :
			x, _ := strconv.ParseFloat(parts[1], 64)
			y, _ := strconv.ParseFloat(parts[2], 64)
			z, _ := strconv.ParseFloat(parts[3], 64)
			vert := vector.V3f{ x, y, z }
			verts = append(verts, vert)
			break;
		case "f" :
			a, _ := strconv.Atoi(strings.Split(parts[1],"/")[0])
			b, _ := strconv.Atoi(strings.Split(parts[2],"/")[0])
			c, _ := strconv.Atoi(strings.Split(parts[3],"/")[0])
			face := Face{ a-1, b-1, c-1 }
			faces = append(faces, face)
			break;
		}
	}

	return &Object{verts, faces}
}

//"https://raw.githubusercontent.com/ssloy/tinyrenderer/master/obj/african_head/african_head.obj"