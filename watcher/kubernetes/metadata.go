package kuberneteswatcher

import (
	"strings"
)

const (
	// METAPREFIX is statusbay prefix annotations
	METAPREFIX = "statusbay.io"
)

// GetMetadataByPrefix will return anitasion values key prefix
func GetMetadataByPrefix(annotations map[string]string, search string) []string {

	values := []string{}
	for key, value := range annotations {
		if strings.HasPrefix(key, search) {
			values = append(values, value)
		}
	}
	return values

}

// GetMetadataOrDefault get metadata from annotatiuons or return a default value
func GetMetadataOrDefault(annotations map[string]string, search string, defaultVal string) string {

	res := GetMetadata(annotations, search)
	if res == "" {
		res = defaultVal
	}
	return res
}

// GetMetadata return specific annotation value
func GetMetadata(annotations map[string]string, search string) string {

	for key, value := range annotations {
		if search == key {
			return value
		}
	}

	var empty string
	return empty
}
