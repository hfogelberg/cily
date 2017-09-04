package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/nfnt/resize"
)

var i = flag.String("i", "", "File in")
var o = flag.String("o", "", "File out. Default is timestamp if left blank")
var w = flag.String("w", "", "Width of output file in px. Default is 700")
var r = flag.Bool("r", false, "Remove original file")

func main() {
	flag.Parse()

	var fileName = ""
	var size uint = 700
	var err error

	if *i == "" {
		log.Println("No input file is given")
		return
	}

	if *o == "" {
		fileName = fmt.Sprint(time.Now().Unix()) + ".jpg"
	} else {
		fileName = *o
	}

	if *w != "" {
		w64, err := strconv.ParseUint(*w, 10, 32)
		if err != nil {
			log.Printf("Error converting width %s to integer. %s", *w, err.Error())
		}
		size = uint(w64)
	}

	fResize, err := os.Open(*i)
	if err != nil {
		log.Printf("Error reducing image size %s\n", err.Error())
		return
	}

	img, err := jpeg.Decode(fResize)
	if err != nil {
		log.Printf("Error decoding image %s\n", err.Error())
		return
	}
	defer fResize.Close()

	m := resize.Resize(size, 0, img, resize.NearestNeighbor)
	out, err := os.Create(fileName)
	if err != nil {
		log.Printf("Error resizing image %s\n", err.Error())
		return
	}
	defer out.Close()
	jpeg.Encode(out, m, nil)

	// Remove uncompressed filed
	if *r == true {
		err = os.Remove(*i)
		if err != nil {
			log.Printf("Error removing uncompressed image %s\n", err.Error())
			return
		}
	}

	fmt.Printf("Image compressed with filename %s\n", fileName)
	return
}
