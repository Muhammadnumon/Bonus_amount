package main

func main() {
}
func bonusAmount(sales [] int) int {
	const percentOfBonus = 5
	const bonusPrice = 10_000
	const percent = 100
	sum := 0
	for _, sale := range sales {
		if sale > bonusPrice {
			result := (sale - bonusPrice) * percentOfBonus / percent
			sum = sum + result
		}
	}
	return sum
}
