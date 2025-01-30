package controllers

func isValidCategory(category string) bool {
	switch category {
		case "maths",
		"physics",
		"chemistry",
		"biology",
		"history",
		"english",
		"computer":
		return true
	}
	return false
}