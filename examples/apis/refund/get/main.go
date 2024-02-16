package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/refund"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	refundClient := refund.NewClient(cfg)

	var paymentID int64 = 12344555
	var refundID int64 = 12344555

	ref, err := refundClient.Get(context.Background(), paymentID, refundID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ref)
}