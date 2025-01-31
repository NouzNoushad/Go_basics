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

func isValidExtensions(extension string) bool {
	switch extension {
		case ".mp4",
		".avi",
		".flv",
		".mkv":
		return true
	}
	return false
}