package main

import (
	"flag"
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
		fileName = time.Now().String() + ".jpg"
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

	// 1. Save file to temp directory
	// file, _, err := r.FormFile("file")
	// if err != nil {
	// 	log.Println(err)
	// 	return "", err
	// }
	// defer file.Close()

	// out, err := os.Create(fileName)
	// if err != nil {
	// 	log.Printf("Error creating file %s\n", err.Error())
	// 	return
	// }
	// defer out.Close()

	// _, err = io.Copy(out, file)
	// if err != nil {
	// 	log.Printf("Error copying filen %s\n", err.Error())
	// 	return
	// }

	log.Printf("File to reduce %s\n", *i)
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

	// name, err := resize(img, size)
	// if err != nil {
	// 	log.Printf("Error resizing image %s\n", err.Error())
	// 	return
	// }

	// // 3. Remove uncompressed filed
	// err = os.Remove(path)
	// if err != nil {
	// 	log.Printf("Error removing uncompressed image %")
	// }

	log.Printf("Image compressed %s\n", fileName)
	return
}

// func resize(img image.Image, width uint, string fileOut) (string, error) {
// 	m := resize.Resize(1000, 0, img, resize.NearestNeighbor)
// 	out, err := os.Create(fileOut)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer out.Close()
// 	jpeg.Encode(out, m, nil)

// 	return fileOut, nil
// }
