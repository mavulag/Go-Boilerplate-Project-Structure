package main

import (
	"os"

	"github.com/mavulag/trilabs/internal/apiservercmd"
	"github.com/sirupsen/logrus"
)

func main() {
	app := apiservercmd.App()

	if err := app.Run(os.Args); err != nil {
		logrus.WithError(err).Fatal("could not run application")
	}
}
