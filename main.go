package main

import (
	"fmt"
	"log"

	aero "github.com/aerospike/aerospike-client-go"
)

// InputData model
type InputData struct {
	Timestamp      int64
	Amount         float64
	UserID         int64
	SellerID       int64
	CreditCardHash string
}

func panicOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	// define a client to connect to
	client, err := aero.NewClient("127.0.0.1", 3000)
	panicOnError(err)

	input := InputData{
		Timestamp:      0,
		Amount:         100,
		UserID:         10,
		SellerID:       11,
		CreditCardHash: "credit-card-hash",
	}

	key, err := aero.NewKey("test", "creditcard", "ID")
	panicOnError(err)

	bin := aero.BinMap{
		"ID":       input.UserID,
		"Time":     input.Timestamp,
		"set_name": input.CreditCardHash,
		"amount":   input.Amount,
		"sellerID": input.SellerID,
	}

	err = client.Put(nil, key, bin)
	panicOnError(err)

	record, err := client.Get(nil, key)
	panicOnError(err)

	fmt.Println(record)
}
