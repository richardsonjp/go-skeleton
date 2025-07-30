package model

import (
	"github.com/shopspring/decimal"
	ct "go-skeleton/internal/model/custom_type"
)

type LocalTime = ct.LocalTime
type JsonArr = ct.JsonArr
type StringArr = ct.StringArr
type Json = ct.Json
type Decimal = ct.Decimal

// NewDecimal converts float64 into Decimal type
func NewDecimal(number float64) Decimal {
	return Decimal(decimal.NewFromFloat(number))
}
