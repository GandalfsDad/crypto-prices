package pricing

import "fmt"

type PriceGetter interface {
	GetPrice(string) float64
}

func PrintPrice(pg PriceGetter, s string) {
	p := pg.GetPrice(s)

	fmt.Printf("%s Price is %f", s, p)
}
