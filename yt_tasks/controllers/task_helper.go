package controllers

func isValidCategory(category string) bool {
	switch category {
	case "ui design",
		"development",
		"testing",
		"deployment",
		"maintenance",
		"research",
		"marketing",
		"content creation":
		return true
	}
	return false
}

func isValidPriority(priority string) bool {
	switch priority {
	case "high",
		"medium",
		"low":
		return true
	}
	return false
}

func isValidStatus(status string) bool {
	switch status {
	case "completed",
		"on review",
		"on hold",
		"in progress":
		return true
	}
	return false
}
