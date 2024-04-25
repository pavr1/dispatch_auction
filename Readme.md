## Overview

This is a small application that gets the best bid based on certain criteria requested by client.

## Quick Start

`main.go` page does nothing on its own, to run this please execute the unit tests by running the following command `go test -v`

## Implementation

To access to this functionality `Auction` struct has to be instantiated with the `pkg.New([]*models.Bid, *logrus.Logger)` function.

## Run

To start the process call `Auction.StartAuction()`

## Improvement areas

- CI/CD implementation
- Restful API exposure
- Mocking Unit tests
