# Crypto  Prices

My aim here is to make a quick command line price querier.
This is really just to help me test my knowledge of go.

Implemented So far 
- BNB Avg Price Query

## Usage
To build:

```cmd
go build -o bin/crypto-prices cmd/main.go
```

To run:
```cmd
bin/crypto-prices {PAIR}
```
This should work for any PAIR that is available on Binance.

