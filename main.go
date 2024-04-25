package main

import (
	"errors"
	"os"

	"github.com/pavr1/dispatch_auction/src/models"
	log "github.com/sirupsen/logrus"
)

func main() {
	initializeLogging()
}

var (
	bestBidKey    string
	bestBidAmount float64
)

func StartAuction(bids map[string]*models.Bid) error {
	if len(bids) == 0 {
		return errors.New("bids are empty")
	}

	setBestBid(bids)

	return nil
}

func processBids(bestBidKey string, bestBidAmount float64, bids map[string]*models.Bid) {
	for k, v := range bids {
		//if the key is the same, then skip
		if k == bestBidKey {
			continue
		}

		if v.CurrentBid < bestBidAmount {
			incrementBid(k, v)
			setBestBid(bids)
		}
	}
}

func incrementBid(bidder string, bid *models.Bid) {
	for {
		if bid.CurrentBid > bestBidAmount {
			return
		}

		if bid.CurrentBid >= bid.MaxBid {
			//out of budget, close bid
			log.WithFields(log.Fields{"Bidder": bidder, "Amount Incremented": bid.CurrentBid, "Best Bid": bestBidAmount}).Info("Bid budget exceeded, closing bid...")
			bid.Status = "closed"
			return
		}

		bid.CurrentBid += bid.BidIncrement

		log.WithFields(log.Fields{"Bidder": bidder, "Amount Incremented": bid.CurrentBid, "Best Bid": bestBidAmount}).Info("Incrementing bid...")
	}
}

func setBestBid(bids map[string]*models.Bid) (string, float64) {
	key := ""
	var bestBid float64

	for k, v := range bids {
		if bestBid < v.CurrentBid {
			key = k
			bestBid = v.CurrentBid
		}
	}

	log.WithFields(log.Fields{"Bidder": key, "Best Bid": bestBid}).Info("Best bid found")

	return key, bestBid
}

func initializeLogging() {
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}
