package kuberneteswatcher

import (
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	//ANNOTATION_PREFIX is statusbay prefix annotations
	ANNOTATION_PREFIX = "statusbay.io"

	//ANNOTATION_PROGRESS_DEADLINE_SECONDS is statusbay progress deadline seconds
	ANNOTATION_PROGRESS_DEADLINE_SECONDS = "progress-deadline-seconds"
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

// GetMetadataOrDefault get metadata from annotations or return a default value
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

//GetMetricsDataFromAnnotations return list of metrics from annotations
func GetMetricsDataFromAnnotations(annotations map[string]string) []Metrics {

	metrics := []Metrics{}
	prefix := fmt.Sprintf("%s/metrics-", ANNOTATION_PREFIX)

	for key, val := range annotations {
		if strings.HasPrefix(key, prefix) {
			metricKey := strings.Replace(key, prefix, "", 1)
			metricData := strings.Split(metricKey, "-")
			if len(metricData) == 0 {
				log.WithFields(log.Fields{
					"key":   key,
					"value": val,
				}).Warn("Invalid annotation metric")
				continue
			}
			metric := Metrics{
				Provider: metricData[0],
				Name:     strings.Join(metricData[1:], " "),
				Query:    val,
			}
			metrics = append(metrics, metric)
		}
	}

	return metrics

}

//GetAlertsDataFromAnnotations return list of alerts from annotations
func GetAlertsDataFromAnnotations(annotations map[string]string) []Alerts {

	alerts := []Alerts{}
	prefix := fmt.Sprintf("%s/alerts-", ANNOTATION_PREFIX)

	for key, val := range annotations {
		if strings.HasPrefix(key, prefix) {
			metricKey := strings.Replace(key, prefix, "", 1)

			alert := Alerts{
				Provider: metricKey,
				Tags:     val,
			}
			alerts = append(alerts, alert)
		}
	}

	return alerts

}

//GetProgressDeadlineApply returns the maximum apply progress. if the field not exists in annotation list default value will returned
func GetProgressDeadlineApply(annotations map[string]string, defaultValue int64) int64 {

	progressDeadLineAnnotations := GetMetadata(annotations, fmt.Sprintf("%s/%s", ANNOTATION_PREFIX, ANNOTATION_PROGRESS_DEADLINE_SECONDS))
	progressDeadLine, err := strconv.ParseInt(progressDeadLineAnnotations, 10, 64)
	if err != nil {
		progressDeadLine = int64(defaultValue)
	}

	return progressDeadLine
}
