package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	var delivery food

	if order == "burger" {
		delivery.preptime = 15
	} else if order == "chips" {
		delivery.preptime = 10
	} else if order == "nuggets" {
		delivery.preptime = 12
	} else {
		return 404
	}
	return delivery.preptime
}
