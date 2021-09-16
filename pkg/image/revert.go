package image

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"log"
	"os"
)

func RevertOneImg(imgpath string) {
	file, err := os.Open(imgpath)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	defer file.Close()
	image, _, err := image.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}
	inverted := imaging.Invert(image)
	imaging.Save(inverted, imgpath)
}
