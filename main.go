package main

import (
	"os"

	"github.com/pavr1/dispatch_auction/src/models"
	"github.com/pavr1/dispatch_auction/src/pkg"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func main() {
	initializeLogging()

	bids := []*models.Bid{
		{
			Name:         "Sasha",
			StartingBid:  50.0,
			CurrentBid:   50.0,
			MaxBid:       80.0,
			BidIncrement: 3.0,
			Status:       "open",
		},
		{
			Name:         "John",
			StartingBid:  60.0,
			CurrentBid:   60.0,
			MaxBid:       82.0,
			BidIncrement: 2.0,
			Status:       "open",
		},
		{
			Name:         "Pat",
			StartingBid:  55.0,
			CurrentBid:   55.0,
			MaxBid:       85.0,
			BidIncrement: 5.0,
			Status:       "open",
		},
	}

	auction := pkg.New(bids, log)
	auction.StartAuction()
	//StartAuction(bids)
}

func initializeLogging() {
	log = logrus.New()
	log.SetOutput(os.Stdout)
	//log.SetReportCaller(true)
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.DebugLevel)
}
