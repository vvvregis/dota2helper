package image

import (
	"github.com/vvvregis/dota2-helper/pkg/config"
	"github.com/vvvregis/dota2-helper/pkg/db"
	"io/ioutil"
	"log"
	"os/exec"
	"runtime"
)

func Main(config config.Config) {
	files, err := ioutil.ReadDir(config.FilePath)
	if err != nil {
		log.Fatalln(err)
	}
	for _, file := range files {
		filepath := config.FilePath + file.Name()
		CropImage(filepath)
		RevertOneImg(filepath)
		id := GetPhotoText(filepath)
		url := "https://ru.dotabuff.com/players/" + id

		if runtime.GOOS == systemLinux {
			err = exec.Command("xdg-open", url).Start()
		} else {
			err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		}

		if config.EnabledDbRecord == 1 {
			db.Insert(config, id)
		}
	}

}
