package controllers

// validate status
func isValidStatus(status string) bool {
	switch status {
	case "pending", "progress", "completed", "canceled":
		return true
	}
	return false
}

// validate category
func isValidCategory(category string) bool {
	switch category {
	case "personal",
		"family",
		"friends",
		"work",
		"date",
		"exercise",
		"sports",
		"shopping",
		"meditation",
		"food",
		"sleep",
		"clean",
		"read",
		"movie",
		"music",
		"gaming":
		return true
	}
	return false
}