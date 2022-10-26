package main

import (
	"math/rand"
	"os"
	"time"

	"world_generator_processor_service/cli"

	"github.com/sirupsen/logrus"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// @title WorldGenerator Processor API
// @version 1.0
// @description Swagger API for WorldGenerator Processor.
// @contact.name Oleh Mushka
// @contact.email olegamysk@gmail.com
// @BasePath /api/
func main() {
	if err := cli.Execute(os.Args[1:]); err != nil {
		logrus.WithField("msg", "cli.Execute error").Error(err)
		os.Exit(1)
	}
}
