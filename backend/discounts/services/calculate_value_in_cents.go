package services

func CalculateValueInCents(priceInCents uint64, percentageDiscount float32) int32 {
	return int32(float32(priceInCents/100) * percentageDiscount)
}