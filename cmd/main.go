package main

import (
	"os"

	"github.com/GandalfsDad/crypto-prices/pkg/pricing"
)

func main() {
	pair := os.Args[1]

	pg := pricing.BnbAvgPriceGetter{}

	pricing.PrintPrice(pg, pair)
}
