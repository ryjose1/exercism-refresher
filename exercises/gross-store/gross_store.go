package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if amount, found := units[unit]; found {
		bill[item] = amount
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	removeAmount, found := units[unit]
	if !found {
		return false
	}
	currentAmount, found := bill[item]
	if !found {
		return false
	}

	if removeAmount > currentAmount {
		return false
	} else if removeAmount == currentAmount {
		delete(bill, item)
	} else {
		bill[item] = currentAmount - removeAmount
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	amount, found := bill[item]
	if !found {
		return 0, false
	}
	return amount, true
}
