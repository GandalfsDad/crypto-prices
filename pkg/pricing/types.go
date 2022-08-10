package pricing

type priceGetter interface {
	getPrice() float64
}
