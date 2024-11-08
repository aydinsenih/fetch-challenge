package receipt

import (
	"math"
	"strings"
	"time"
)

type Receipt struct {
	Retailer     string  `json:"retailer" validate:"required"`
	PurchaseDate string  `json:"purchaseDate" validate:"required"`
	PurchaseTime string  `json:"purchaseTime" validate:"required"`
	Items        []Items `json:"items" validate:"required"`
	Total        float64 `json:"total,string" validate:"required"`
}

type Items struct {
	ShortDescription string  `json:"shortDescription" validate:"required"`
	Price            float64 `json:"price,string" validate:"required"`
}

func (r Receipt) Calc() int {

	return findAlphaNumericPoint(r.Retailer) +
		isTotalRound(r.Total) +
		isTotalMultipleOf25(r.Total) +
		findPair(r.Items) +
		itemDescCalc(r.Items) +
		isDateOdd(r.PurchaseDate) +
		between2and4(r.PurchaseTime)
}

func findAlphaNumericPoint(s string) int {
	count := 0
	for _, c := range s {
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			count++
		}
	}
	return count
}

func isTotalRound(t float64) int {
	if math.Mod(t, 1) == 0 {
		return 50
	}
	return 0
}

func isTotalMultipleOf25(t float64) int {
	m := int(math.Mod(t, 1) * 100)
	if m%25 == 0 {
		return 25
	}
	return 0
}

func findPair(l []Items) int {
	len := len(l)
	return int(math.Floor(float64(len)/2)) * 5
}

func itemDescCalc(d []Items) int {
	result := 0
	for _, item := range d {
		l := len(strings.Trim(item.ShortDescription, " "))
		if l%3 == 0 {
			result = result + int(math.Ceil(item.Price*0.2))
		}
	}
	return result
}

func isDateOdd(d string) int {
	date, err := time.Parse("2006-01-02", d)
	if err != nil {
		return 0
	}
	return date.Day() % 2 * 6
}

func between2and4(d string) int {
	t, err := time.Parse("15:04", d)
	if err != nil {
		return 0
	}
	if t.Hour() == 14 && t.Minute() != 0 {
		return 10
	}
	if t.Hour() == 15 {
		return 10
	}
	return 0
}
