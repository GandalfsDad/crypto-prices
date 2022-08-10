package pricing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type BnbAvgPriceGetter struct{}

type BnBAvgReturn struct {
	Mins  int32  `json: "mins"`
	Price string `json: "price"`
}

var bnbAvgPriceAddress = "https://api.binance.com/api/v3/avgPrice?symbol="

func (bap BnbAvgPriceGetter) GetPrice(s string) float64 {
	r, err := http.Get(bnbAvgPriceAddress + s)

	if err != nil {
		fmt.Println("Error Getting Response:", err)
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Error Reading Body:", err)
		os.Exit(1)
	}

	data := BnBAvgReturn{}

	err = json.Unmarshal(b, &data)

	if err != nil {
		fmt.Println("Error with response format:", err)
	}

	p, err := strconv.ParseFloat(data.Price, 64)

	if err != nil {
		fmt.Println("Error Converting:", err)
		os.Exit(1)
	}

	return p
}
