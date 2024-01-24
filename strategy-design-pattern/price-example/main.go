package main

import "fmt"

type PriceInterface interface {
	getPrice(originPrice float32) float32
}

type preOrder struct {
}

func (po *preOrder) getPrice(originPrice float32) float32 {
	return originPrice * 0.8
}

type Promotion struct{}

func (pr *Promotion) getPrice(originPrice float32) float32 {
	return originPrice * 0.5
}

type BlackFriday struct{}

func (bf *BlackFriday) getPrice(originPrice float32) float32 {
	return originPrice * 0.3
}

func main() {
	po := preOrder{}
	pr := Promotion{}
	bf := BlackFriday{}

	getPriceStrategies := make(map[string]PriceInterface)
	getPriceStrategies["preOrder"] = &po
	getPriceStrategies["Promotion"] = &pr
	getPriceStrategies["BlackFriday"] = &bf

	//define type Promotion in here
	typePromotion := "BlackFriday"
	origin := 200
	finalPrice := getPriceStrategies[typePromotion].getPrice(float32(origin))

	fmt.Println("Price ", typePromotion, "::", finalPrice)
}
