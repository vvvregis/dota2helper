package image

import (
	"bufio"
	"log"
	"os"
)

func ReadFile() string {

	file, err := os.Open("output.txt")

	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	var steamId string

	for fileScanner.Scan() {
		steamId = fileScanner.Text()
		break
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return steamId
}
