// Copyright 2020 Business Process Technologies. All rights reserved.

package main

import (
	"fmt"
	"math"
)

func GetFee(feePerc int) {
	payment := 100000000

	feeReal := float64(feePerc) / 100.0
	PaymentPerc := float64(payment) / 100.0
	feeRaw := PaymentPerc * feeReal
	feeSysRewardRaw := feeRaw / 100.0 * 50.0
	feeRefRewardRaw := feeRaw / 100.0 * 25.0

	fee := int64(math.Round(feeRaw))
	feeSysReward := int64(math.Round(feeSysRewardRaw))
	feeRefReward := int64(math.Round(feeRefRewardRaw))
	feeBurn := fee - (feeSysReward + feeRefReward)

	fmt.Printf("payment %v feePerc: %v feeRaw: %f fee: %v Sys:%v Ref:%v Burn:%v \n", payment, feePerc, feeRaw, fee, feeSysReward, feeRefReward, feeBurn)
}

func main() {

	fees := []int{50, 100, 150, 200}

	for _, fee := range fees {
		GetFee(fee)
	}
}
