package handlers

import (
	"bufio"
	"log"
	"os"

	configurations "github.com/golang-line-ingestor/configurations"
	models "github.com/golang-line-ingestor/models"
	logger "github.com/sirupsen/logrus"
	resty "gopkg.in/resty.v1"
)

func Processor() {
	file, err := os.Open(configurations.LineFileFullpath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data := models.SDLRequest{Data: scanner.Text()}
		_, err := resty.R().
			SetBody(data).
			Post(configurations.SDLEndpoint)
		if err != nil {
			logger.Error(data)
		}
	}
}
