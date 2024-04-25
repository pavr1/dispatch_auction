package main

import (
	"github.com/pavr1/dispatch_auction/src/pkg"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func main() {
	log = pkg.InitializeLogging()
}
