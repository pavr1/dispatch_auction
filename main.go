package main

import (
	"errors"
	"os"

	"github.com/pavr1/dispatch_auction/src/models"
	log "github.com/sirupsen/logrus"
)

func main() {
	initializeLogging()

	bids := map[string]*models.Bid{}

	bids["Sasha"] = &models.Bid{
		StartingBid:  50.0,
		CurrentBid:   50.0,
		MaxBid:       60.0,
		BidIncrement: 3.0,
		Status:       "open",
	}
	bids["John"] = &models.Bid{
		StartingBid:  60.0,
		CurrentBid:   60.0,
		MaxBid:       82.0,
		BidIncrement: 2.0,
		Status:       "open",
	}
	bids["Pat"] = &models.Bid{
		StartingBid:  55.0,
		CurrentBid:   55.0,
		MaxBid:       85.0,
		BidIncrement: 5.0,
		Status:       "open",
	}

	StartAuction(bids)
}

func StartAuction(bids map[string]*models.Bid) error {
	if len(bids) == 0 {
		return errors.New("bids are empty")
	}

	bestBidKey, bestBidAmount := setBestBid(bids)
	processBids(bestBidKey, bestBidAmount, bids)

	return nil
}

func processBids(bestBidKey *string, bestBidAmount *float64, bids map[string]*models.Bid) {
	log.WithFields(log.Fields{"Best Bid": bestBidAmount}).Info("Processing bids...")

	for k, v := range bids {
		//if the key is the same, then skip
		if k == *bestBidKey {
			continue
		}

		//if the bid is closed then skip
		if v.Status == "closed" {
			continue
		}

		log.WithFields(log.Fields{"Bidder": k, "Current Amount": v.CurrentBid, "Best Bid": bestBidAmount}).Info("Processing bid!")

		if v.CurrentBid < *bestBidAmount {
			log.WithFields(log.Fields{"Bidder": k, "Current Amount": v.CurrentBid, "Best Bid": bestBidAmount}).Info("Bid below offer, increasing...")

			incrementBid(k, bestBidKey, bestBidAmount, v)
		}
	}
}

func incrementBid(bidder string, bestBidKey *string, bestBidAmount *float64, bid *models.Bid) {
	for {
		if bid.CurrentBid > *bestBidAmount {
			log.WithFields(log.Fields{"Bidder": bidder, "Amount Incremented": bid.CurrentBid, "Best Bid": bestBidAmount}).Info("Bid Incremented!")

			bestBidKey = &bidder
			bestBidAmount = &bid.CurrentBid

			return
		}

		if bid.CurrentBid >= bid.MaxBid {
			//out of budget, close bid
			log.WithFields(log.Fields{"Bidder": bidder, "Amount Incremented": bid.CurrentBid, "Best Bid": bestBidAmount}).Warn("Bid budget exceeded, closing bid!")
			bid.Status = "closed"
			return
		}

		bid.CurrentBid += bid.BidIncrement

		//log.WithFields(log.Fields{"Bidder": bidder, "Amount Incremented": bid.CurrentBid, "Best Bid": bestBidAmount}).Info("Incrementing bid!")
	}
}

func setBestBid(bids map[string]*models.Bid) (*string, *float64) {
	key := ""
	var bestBid float64

	for k, v := range bids {
		if v.Status == "closed" {
			continue
		}

		if bestBid < v.CurrentBid {
			key = k
			bestBid = v.CurrentBid
		}
	}

	log.WithFields(log.Fields{"Bidder": key, "Best Bid": bestBid}).Info("Best bid found")

	return &key, &bestBid
}

func initializeLogging() {
	log.SetOutput(os.Stdout)
	//log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}
