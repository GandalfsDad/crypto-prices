package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	pair := os.Args[1]
	s, err := http.Get("https://api.binance.com/api/v3/avgPrice?symbol=" + pair)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, s.Body)
}
