package main

import (
	"flag"
	"image"
	"log"
	"os"

	"github.com/signintech/gopdf"
)

func main() {

	var imagePath string
	flag.StringVar(&imagePath, "path", "bar", "a string var")
	flag.Parse()

	w, h := getImageDimension(imagePath)
	startW := (float64)(w) / (4.0)
	startH := (float64)(h) / (4.0)

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: (float64)(w), H: (float64)(h)}})
	pdf.AddPage()

	// use image holder by []byte
	imgH1, err := gopdf.ImageHolderByBytes(getImageBytes(imagePath))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	pdf.ImageByHolder(imgH1, startW, startH, nil)

	pdf.WritePdf("image.pdf")
}

func getImageBytes(imagePath string) []byte {
	img, err := os.ReadFile(imagePath)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return img
}

func getImageDimension(imagePath string) (int, int) {
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal(err.Error())
		return -1, -1
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		log.Fatal(err.Error())
		return -1, -1
	}

	return image.Width, image.Height
}
