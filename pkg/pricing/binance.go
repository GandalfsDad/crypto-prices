package pricing

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type bnbAvgPriceGetter struct{}

var bnbAvgPriceAddress = "https://api.binance.com/api/v3/avgPrice?symbol="

func (bap bnbAvgPriceGetter) getPrice(s string) float64 {
	resp, err := http.Get(bnbAvgPriceAddress + s)

	if err != nil {
		fmt.Println("Error:", err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Errors:", err)
	}

	conv, err := strconv.ParseFloat(strings.Split(strings.Split(string(body), ",")[1], ":")[1], 64)

	return conv
}
