package main

import (
	"github.com/vvvregis/dota2-helper/pkg/config"
	"github.com/vvvregis/dota2-helper/pkg/image"
	"log"
)

func main() {
	conf, err := config.LoadConfiguration("pkg/config/image/config.json")
	if err != nil {
		log.Fatalln(err)
	}

	image.Main(conf)
}
