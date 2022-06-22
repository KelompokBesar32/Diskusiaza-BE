package library

import "strings"

func CheckExtensionForImage(extension string) bool {
	allowed := []string{
		".png", ".jpg", ".jpeg",
	}
	extension = strings.ToLower(extension)
	for _, v := range allowed {
		if v == extension {
			return true
		}
	}
	return false
}
