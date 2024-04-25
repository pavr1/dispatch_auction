package pkg

import (
	"errors"

	"github.com/pavr1/dispatch_auction/src/models"
	"github.com/sirupsen/logrus"
)

type Auction struct {
	bestBidKey    *string
	bestBidAmount *float64
	Bids          []*models.Bid
	log           *logrus.Logger
}

func New(bids []*models.Bid, log *logrus.Logger) *Auction {
	return &Auction{
		Bids: bids,
		log:  log,
	}
}

func (a *Auction) StartAuction() error {
	if len(a.Bids) == 0 {
		return errors.New("bids are empty")
	}

	name, bestBid := a.getBestBid()
	a.bestBidKey = &name
	a.bestBidAmount = &bestBid

	bid := a.processBids()

	a.log.WithFields(logrus.Fields{"Name": bid.Name, "Amount": bid.CurrentBid}).Info("WINNING BID!!!")

	return nil
}

func (a *Auction) processBids() *models.Bid {
	if len(a.Bids) == 1 {
		return a.Bids[0]
	}

	for _, v := range a.Bids {
		//if the key is the same, then skip
		if v.Name == *a.bestBidKey {
			continue
		}

		//if the bid is closed then skip
		if v.Status == "closed" {
			continue
		}

		//log.WithFields(log.Fields{"Name": v.Name, "Current Amount": v.CurrentBid, "Best Bid": bestBidAmount}).Info("Processing bid!")

		if v.CurrentBid < *a.bestBidAmount {
			//a.log.WithFields(logrus.Fields{"Name": v.Name, "Current Amount": v.CurrentBid, "Best Bid": a.bestBidAmount}).Info("Bid below offer, increasing...")

			a.incrementBid(v.Name, v)
		}

		name, bestBid := a.getBestBid()
		a.bestBidKey = &name
		a.bestBidAmount = &bestBid
	}

	a.Bids = a.cleanUpClosedBids(a.Bids)

	return a.processBids()
}

func (a *Auction) incrementBid(bidder string, bid *models.Bid) {
	for {
		if bid.CurrentBid > *a.bestBidAmount {
			a.log.WithFields(logrus.Fields{"Name": bidder, "Bid Incremented": bid.CurrentBid, "Best Bid": a.bestBidAmount, "Status": bid.Status}).Info("Bid Incremented!")

			return
		}

		if bid.CurrentBid >= bid.MaxBid {
			//out of budget, close bid
			a.log.WithFields(logrus.Fields{"Name": bidder, "Amount Incremented": bid.CurrentBid, "Best Bid": a.bestBidAmount}).Warn("Bid budget exceeded, closing bid!")
			bid.Status = "closed"
			return
		}

		bid.CurrentBid += bid.BidIncrement

		//log.WithFields(log.Fields{"Name": bidder, "Amount Incremented": bid.CurrentBid, "Best Bid": bestBidAmount}).Info("Incrementing bid!")
	}
}

func (a *Auction) cleanUpClosedBids(bids []*models.Bid) []*models.Bid {
	for i := len(bids) - 1; i >= 0; i-- {
		bid := bids[i]

		if bid.Status == "closed" {
			bids = a.deleteElement(bids, i)
		}
	}

	return bids
}

func (a *Auction) deleteElement(slice []*models.Bid, index int) []*models.Bid {
	return append(slice[:index], slice[index+1:]...)
}

func (a *Auction) getBestBid() (string, float64) {
	name := ""
	var bestBid float64

	for _, v := range a.Bids {
		if v.Status == "closed" {
			continue
		}

		if bestBid < v.CurrentBid {
			name = v.Name
			bestBid = v.CurrentBid
		}
	}

	a.log.WithFields(logrus.Fields{"Name": name, "Best Bid": bestBid}).Info("Best bid found")

	return name, bestBid
}
