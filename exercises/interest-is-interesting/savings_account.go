package savings

import "math"

// InterestRate calculates the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		return -3.213
	case balance < 1000:
		return 0.5
	case balance < 5000:
		return 1.621
	case balance >= 5000:
		return 2.475
	default:
		return 0
	}
}

// InterestRate calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return math.Copysign(balance*float64(InterestRate(balance))/100.0, balance)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) (years int) {
	if Interest(balance) < 0 && balance < targetBalance {
		panic("Balance will never reach target due to negative growth")
	}

	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		years += 1
	}
	return
}
