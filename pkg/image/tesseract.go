package image

import (
	"github.com/otiai10/gosseract"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const systemLinux = "linux"

func GetPhotoText(imgpath string) string {
	var recognizedText string
	currentSystem := runtime.GOOS

	if currentSystem == systemLinux {
		recognizedText = recognizeLinux(imgpath)
	} else {
		recognizedText = recognizeWindows(imgpath)
	}

	os.Remove(imgpath)

	return recognizedText
}

func recognizeLinux(imgpath string) string {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage(imgpath)
	text, err := client.Text()
	if err != nil {
		log.Fatalln(err)
	}

	return text
}

func recognizeWindows(imgpath string) string {
	commandArr := make([]string, 3, 6)
	commandArr[0] = "tesseract "
	commandArr[1] = imgpath
	commandArr[2] = " output"
	commandStr := strings.Join(commandArr, "")

	cmd := exec.Command("cmd.exe", "/C", commandStr)

	_, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalln(err)
	}

	text := ReadFile()

	return text

}
