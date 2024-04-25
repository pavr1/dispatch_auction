package models

type Bid struct {
	Name         string
	StartingBid  float64 "json:`starting_bid`"
	MaxBid       float64 "json:`max_bid`"
	BidIncrement float64 "json:`bad_increment`"
	CurrentBid   float64
	Status       string //Open, close
}
