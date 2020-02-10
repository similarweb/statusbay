package kuberneteswatcher

import (
	"fmt"
	"testing"
)

func TestGetMetadataByPrefix(t *testing.T) {
	annotations := map[string]string{
		"foo":   "foo-val",
		"foo-2": "foo-val-2",
		"bar":   "bar-val",
	}
	result := GetMetadataByPrefix(annotations, "foo")

	if len(result) != 2 {
		t.Fatalf("unexpected annotation values count, got %d expected %d", len(result), 2)
	}

}
func TestGetMetadataOrDefault(t *testing.T) {
	annotations := map[string]string{
		"foo":   "foo-val",
		"foo-2": "foo-val-2",
		"bar":   "bar-val",
	}

	t.Run("get_annotation_value", func(t *testing.T) {
		result := GetMetadataOrDefault(annotations, "foo", "val")

		if result != "foo-val" {
			t.Fatalf("unexpected annotation value, got %s expected %s", result, "foo-val")
		}
	})

	t.Run("get_annotation_value_with_default", func(t *testing.T) {
		result := GetMetadataOrDefault(annotations, "not_exists", "default value")

		if result != "default value" {
			t.Fatalf("unexpected annotation value, got %s expected %s", result, "default value")
		}
	})

}

func TestGetMetadata(t *testing.T) {
	annotations := map[string]string{
		"foo":   "foo-val",
		"foo-2": "foo-val-2",
		"bar":   "bar-val",
	}

	t.Run("get_annotation_value", func(t *testing.T) {
		result := GetMetadata(annotations, "foo")

		if result != "foo-val" {
			t.Fatalf("unexpected annotation value, got %s expected %s", result, "foo-val")
		}
	})

	t.Run("get_un_exists_annotation", func(t *testing.T) {
		result := GetMetadata(annotations, "not_exists")

		if result != "" {
			t.Fatalf("unexpected annotation value, got %s expected %s", result, "")
		}
	})

}

func TestGetMetricsDataFromAnnotations(t *testing.T) {

	annotations := map[string]string{
		fmt.Sprintf("%s/metrics-providername-metric-name", ANNOTATION_PREFIX): "metric query 1",
	}

	metrics := GetMetricsDataFromAnnotations(annotations)

	if len(metrics) != 1 {
		t.Fatalf("unexpected annotation metrics count, got %d expected %d", len(metrics), 1)
	}

	metric := metrics[0]

	if metric.Name != "metric name" {
		t.Fatalf("unexpected metric name, got %s expected %s", metric.Name, "metric name")
	}

	if metric.Provider != "providername" {
		t.Fatalf("unexpected metric provider, got %s expected %s", metric.Provider, "providername")
	}

}

func TestGetAlertsDataFromAnnotations(t *testing.T) {
	annotations := map[string]string{
		fmt.Sprintf("%s/alerts-providername", ANNOTATION_PREFIX): "foo,foo1",
	}

	alerts := GetAlertsDataFromAnnotations(annotations)

	if len(alerts) != 1 {
		t.Fatalf("unexpected annotation alerts count, got %d expected %d", len(alerts), 1)
	}

	alert := alerts[0]

	if alert.Provider != "providername" {
		t.Fatalf("unexpected alert provider, got %s expected %s", alert.Provider, "providername")
	}
	if alert.Tags != "foo,foo1" {
		t.Fatalf("unexpected alert tags, got %s expected %s", alert.Tags, "foo,foo1")
	}
}

func TestGetProgressDeadlineApply(t *testing.T) {

	t.Run("get_progress_deadline_annotation", func(t *testing.T) {

		annotations := map[string]string{
			fmt.Sprintf("%s/%s", ANNOTATION_PREFIX, ANNOTATION_PROGRESS_DEADLINE_SECONDS): "10",
		}
		progressDeadlineSeconds := GetProgressDeadlineApply(annotations, 2)

		if progressDeadlineSeconds != 10 {
			t.Fatalf("unexpected %s annotation value, got %d expected %d", ANNOTATION_PROGRESS_DEADLINE_SECONDS, progressDeadlineSeconds, 10)

		}
	})
	t.Run("get_un_exists_progress_deadline_annotation", func(t *testing.T) {

		annotations := map[string]string{}
		progressDeadlineSeconds := GetProgressDeadlineApply(annotations, 60)

		if progressDeadlineSeconds != 60 {
			t.Fatalf("unexpected %s annotation value, got %d expected %d", ANNOTATION_PROGRESS_DEADLINE_SECONDS, progressDeadlineSeconds, 60)

		}
	})

}
