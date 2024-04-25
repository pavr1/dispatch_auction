package main

import (
	"errors"
	"os"

	"github.com/pavr1/dispatch_auction/src/models"
	log "github.com/sirupsen/logrus"
)

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

	StartAuction(bids)
}

func StartAuction(bids []*models.Bid) error {
	if len(bids) == 0 {
		return errors.New("bids are empty")
	}

	bestBidKey, bestBidAmount := setBestBid(bids)
	bid := processBids(bestBidKey, bestBidAmount, bids)

	log.WithFields(log.Fields{"Bidder": bid.Name, "Amount": bid.CurrentBid}).Info("WINNING BID!!!")

	return nil
}

func processBids(bestBidKey string, bestBidAmount float64, bids []*models.Bid) *models.Bid {
	if len(bids) == 1 {
		return bids[0]
	}

	for _, v := range bids {
		//if the key is the same, then skip
		if v.Name == bestBidKey {
			continue
		}

		//if the bid is closed then skip
		if v.Status == "closed" {
			continue
		}

		//log.WithFields(log.Fields{"Bidder": v.Name, "Current Amount": v.CurrentBid, "Best Bid": bestBidAmount}).Info("Processing bid!")

		if v.CurrentBid < bestBidAmount {
			log.WithFields(log.Fields{"Bidder": v.Name, "Current Amount": v.CurrentBid, "Best Bid": bestBidAmount}).Info("Bid below offer, increasing...")

			incrementBid(v.Name, bestBidKey, bestBidAmount, v)
		}

		bestBidKey, bestBidAmount = setBestBid(bids)
	}

	bids = cleanUpClosedBids(bids)

	return processBids(bestBidKey, bestBidAmount, bids)
}

func cleanUpClosedBids(bids []*models.Bid) []*models.Bid {
	for i := len(bids) - 1; i >= 0; i-- {
		bid := bids[i]

		if bid.Status == "closed" {
			bids = deleteElement(bids, i)
		}
	}

	return bids
}

func deleteElement(slice []*models.Bid, index int) []*models.Bid {
	return append(slice[:index], slice[index+1:]...)
}

func incrementBid(bidder string, bestBidKey string, bestBidAmount float64, bid *models.Bid) {
	for {
		if bid.CurrentBid > bestBidAmount {
			log.WithFields(log.Fields{"Bidder": bidder, "Amount Incremented": bid.CurrentBid, "Best Bid": bestBidAmount, "Status": bid.Status}).Info("Bid Incremented!")

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

func setBestBid(bids []*models.Bid) (string, float64) {
	key := ""
	var bestBid float64

	for _, v := range bids {
		if v.Status == "closed" {
			continue
		}

		if bestBid < v.CurrentBid {
			key = v.Name
			bestBid = v.CurrentBid
		}
	}

	log.WithFields(log.Fields{"Bidder": key, "Best Bid": bestBid}).Info("Best bid found")

	return key, bestBid
}

func initializeLogging() {
	log.SetOutput(os.Stdout)
	//log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}
