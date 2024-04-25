package pkg

import (
	"testing"

	"github.com/pavr1/dispatch_auction/src/models"
)

// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'
type Tester struct {
	Bids           []*models.Bid
	ExpectedName   string
	ExpectedAmount float64
}

func Test_StartAuction(t *testing.T) {
	log := InitializeLogging()

	tests := loadStartAuctionBids()

	for _, test := range tests {
		auction := New(test.Bids, log)
		result, err := auction.StartAuction()
		if err != nil {
			t.Errorf("Unexpected error found: %v", err)
		}

		if result.Name != test.ExpectedName {
			t.Errorf("Expected name '%s' but got '%s'", test.ExpectedName, result.Name)
		}

		if result.CurrentBid != test.ExpectedAmount {
			t.Errorf("Expected bid '%f' but got '%f'", test.ExpectedAmount, result.CurrentBid)
		}
	}
}

func Test_getBestBid(t *testing.T) {
	log := InitializeLogging()

	tests := loadgetBestBidBids()

	for _, test := range tests {
		auction := New(test.Bids, log)
		name, bestBib := auction.getBestBid()

		if name != test.ExpectedName {
			t.Errorf("Expected name '%s' but got '%s'", test.ExpectedName, name)
		}

		if bestBib != test.ExpectedAmount {
			t.Errorf("Expected bid '%f' but got '%f'", test.ExpectedAmount, bestBib)
		}
	}
}

func Test_cleanUpClosedBids(t *testing.T) {
	log := InitializeLogging()

	tests := loadcleanUpClosedBids()

	for _, test := range tests {
		auction := New(test.Bids, log)
		result := auction.cleanUpClosedBids(auction.Bids)

		if len(result) != int(test.ExpectedAmount) {
			t.Errorf("Expected clean slice length '%d' but got '%d'", int(test.ExpectedAmount), len(result))
		}
	}
}

func Test_incrementBid(t *testing.T) {
	log := InitializeLogging()

	tests := loadincrementBids()

	for _, test := range tests {
		auction := New(test.Bids, log)
		name, bestBid := auction.getBestBid()
		auction.bestBidKey = &name
		auction.bestBidAmount = &bestBid

		bid := auction.Bids[0]
		auction.incrementBid("", bid)

		if bid.Name != test.ExpectedName {
			t.Errorf("Expected name '%s' but got '%s'", test.ExpectedName, bid.Name)
		}

		if bid.CurrentBid != test.ExpectedAmount {
			t.Errorf("Expected bid '%f' but got '%f'", test.ExpectedAmount, bid.CurrentBid)
		}
	}
}

func loadincrementBids() []Tester {
	return []Tester{
		{
			Bids: []*models.Bid{
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
			},
			ExpectedName:   "Sasha",
			ExpectedAmount: 62.0,
		},
		{
			Bids: []*models.Bid{
				{
					Name:         "Riley",
					StartingBid:  700.0,
					CurrentBid:   700.0,
					MaxBid:       725.0,
					BidIncrement: 2.0,
					Status:       "open",
				},
				{
					Name:         "Morgan",
					StartingBid:  599.0,
					CurrentBid:   599.0,
					MaxBid:       725.0,
					BidIncrement: 15.0,
					Status:       "open",
				},
				{
					Name:         "Charlie",
					StartingBid:  625.0,
					CurrentBid:   625.0,
					MaxBid:       725.0,
					BidIncrement: 8.0,
					Status:       "open",
				},
			},
			ExpectedName:   "Riley",
			ExpectedAmount: 702.0,
		},
		{
			Bids: []*models.Bid{
				{
					Name:         "Alex",
					StartingBid:  2500.0,
					CurrentBid:   2500.0,
					MaxBid:       3000.0,
					BidIncrement: 500.0,
					Status:       "open",
				},
				{
					Name:         "Jesse",
					StartingBid:  2800.0,
					CurrentBid:   2800.0,
					MaxBid:       3100.0,
					BidIncrement: 201.0,
					Status:       "open",
				},
				{
					Name:         "Drew",
					StartingBid:  2501.0,
					CurrentBid:   2501.0,
					MaxBid:       3200.0,
					BidIncrement: 247.0,
					Status:       "open",
				},
			},
			ExpectedName:   "Alex",
			ExpectedAmount: 3000.0,
		},
	}
}

func loadgetBestBidBids() []Tester {
	return []Tester{
		{
			Bids: []*models.Bid{
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
			},
			ExpectedName:   "John",
			ExpectedAmount: 60.0,
		},
		{
			Bids: []*models.Bid{
				{
					Name:         "Riley",
					StartingBid:  700.0,
					CurrentBid:   700.0,
					MaxBid:       725.0,
					BidIncrement: 2.0,
					Status:       "open",
				},
				{
					Name:         "Morgan",
					StartingBid:  599.0,
					CurrentBid:   599.0,
					MaxBid:       725.0,
					BidIncrement: 15.0,
					Status:       "open",
				},
				{
					Name:         "Charlie",
					StartingBid:  625.0,
					CurrentBid:   625.0,
					MaxBid:       725.0,
					BidIncrement: 8.0,
					Status:       "open",
				},
			},
			ExpectedName:   "Riley",
			ExpectedAmount: 700.0,
		},
		{
			Bids: []*models.Bid{
				{
					Name:         "Alex",
					StartingBid:  2500.0,
					CurrentBid:   2500.0,
					MaxBid:       3000.0,
					BidIncrement: 500.0,
					Status:       "open",
				},
				{
					Name:         "Jesse",
					StartingBid:  2800.0,
					CurrentBid:   2800.0,
					MaxBid:       3100.0,
					BidIncrement: 201.0,
					Status:       "open",
				},
				{
					Name:         "Drew",
					StartingBid:  2501.0,
					CurrentBid:   2501.0,
					MaxBid:       3200.0,
					BidIncrement: 247.0,
					Status:       "open",
				},
			},
			ExpectedName:   "Jesse",
			ExpectedAmount: 2800.0,
		},
	}
}

func loadcleanUpClosedBids() []Tester {
	return []Tester{
		{
			Bids: []*models.Bid{
				{
					Name:         "Sasha",
					StartingBid:  50.0,
					CurrentBid:   50.0,
					MaxBid:       80.0,
					BidIncrement: 3.0,
					Status:       "closed",
				},
				{
					Name:         "John",
					StartingBid:  60.0,
					CurrentBid:   60.0,
					MaxBid:       82.0,
					BidIncrement: 2.0,
					Status:       "closed",
				},
				{
					Name:         "Pat",
					StartingBid:  55.0,
					CurrentBid:   55.0,
					MaxBid:       85.0,
					BidIncrement: 5.0,
					Status:       "open",
				},
			},
			ExpectedName: "",
			//this is the expected total count of the slice after cleanup
			ExpectedAmount: 1,
		},
		{
			Bids: []*models.Bid{
				{
					Name:         "Riley",
					StartingBid:  700.0,
					CurrentBid:   700.0,
					MaxBid:       725.0,
					BidIncrement: 2.0,
					Status:       "open",
				},
				{
					Name:         "Morgan",
					StartingBid:  599.0,
					CurrentBid:   599.0,
					MaxBid:       725.0,
					BidIncrement: 15.0,
					Status:       "open",
				},
				{
					Name:         "Charlie",
					StartingBid:  625.0,
					CurrentBid:   625.0,
					MaxBid:       725.0,
					BidIncrement: 8.0,
					Status:       "closed",
				},
			},
			ExpectedName: "",
			//this is the expected total count of the slice after cleanup
			ExpectedAmount: 2,
		},
		{
			Bids: []*models.Bid{
				{
					Name:         "Alex",
					StartingBid:  2500.0,
					CurrentBid:   2500.0,
					MaxBid:       3000.0,
					BidIncrement: 500.0,
					Status:       "open",
				},
				{
					Name:         "Jesse",
					StartingBid:  2800.0,
					CurrentBid:   2800.0,
					MaxBid:       3100.0,
					BidIncrement: 201.0,
					Status:       "open",
				},
				{
					Name:         "Drew",
					StartingBid:  2501.0,
					CurrentBid:   2501.0,
					MaxBid:       3200.0,
					BidIncrement: 247.0,
					Status:       "open",
				},
			},
			ExpectedName:   "",
			ExpectedAmount: 3,
		},
	}
}

func loadStartAuctionBids() []Tester {
	return []Tester{
		{
			Bids: []*models.Bid{
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
			},
			ExpectedName:   "Pat",
			ExpectedAmount: 85.0,
		},
		{
			Bids: []*models.Bid{
				{
					Name:         "Riley",
					StartingBid:  700.0,
					CurrentBid:   700.0,
					MaxBid:       725.0,
					BidIncrement: 2.0,
					Status:       "open",
				},
				{
					Name:         "Morgan",
					StartingBid:  599.0,
					CurrentBid:   599.0,
					MaxBid:       725.0,
					BidIncrement: 15.0,
					Status:       "open",
				},
				{
					Name:         "Charlie",
					StartingBid:  625.0,
					CurrentBid:   625.0,
					MaxBid:       725.0,
					BidIncrement: 8.0,
					Status:       "open",
				},
			},
			ExpectedName:   "Morgan",
			ExpectedAmount: 734.0,
		},
		{
			Bids: []*models.Bid{
				{
					Name:         "Alex",
					StartingBid:  2500.0,
					CurrentBid:   2500.0,
					MaxBid:       3000.0,
					BidIncrement: 500.0,
					Status:       "open",
				},
				{
					Name:         "Jesse",
					StartingBid:  2800.0,
					CurrentBid:   2800.0,
					MaxBid:       3100.0,
					BidIncrement: 201.0,
					Status:       "open",
				},
				{
					Name:         "Drew",
					StartingBid:  2501.0,
					CurrentBid:   2501.0,
					MaxBid:       3200.0,
					BidIncrement: 247.0,
					Status:       "open",
				},
			},
			ExpectedName:   "Drew",
			ExpectedAmount: 3242.0,
		},
	}
}
