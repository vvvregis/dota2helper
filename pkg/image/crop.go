package image

import (
	"fmt"
	"image"
	"log"
	"os"

	"image/png"
)

func CropImage(imgpath string) {
	if err := run(imgpath); err != nil {
		log.Fatalln(err)
	}
}

func run(imgpath string) error {
	img, err := readImage(imgpath)
	if err != nil {
		return err
	}
	img, err = cropImage(img, image.Rect(590, 220, 690, 250))
	if err != nil {
		return err
	}

	return writeImage(img, imgpath)
}

func readImage(imgpath string) (image.Image, error) {
	fd, err := os.Open(imgpath)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	img, _, err := image.Decode(fd)

	if err != nil {
		return nil, err
	}

	return img, nil
}

func cropImage(img image.Image, crop image.Rectangle) (image.Image, error) {
	type subImager interface {
		SubImage(r image.Rectangle) image.Image
	}

	simg, ok := img.(subImager)
	if !ok {
		return nil, fmt.Errorf("image does not support cropping")
	}

	return simg.SubImage(crop), nil
}

func writeImage(img image.Image, name string) error {
	fd, err := os.Create(name)
	if err != nil {
		return err
	}
	defer fd.Close()

	return png.Encode(fd, img)
}
