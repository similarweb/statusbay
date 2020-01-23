package kubernetes_test

import (
	"fmt"
	"net/http"
	"reflect"
	"statusbay/api/kubernetes"
	"strconv"
	"strings"
	"testing"
)

func TestFilterApplication(t *testing.T) {

	t.Run("defaults", func(t *testing.T) {
		expectedFilters := kubernetes.FilterApplications{
			Offset:        0,
			Limit:         20,
			Clusters:      []string{""},
			Namespaces:    []string{""},
			Statuses:      []string{""},
			Application:   "",
			SortBy:        "time",
			SortDirection: "desc",
			From:          0,
			To:            0,
			Distinct:      false,
		}
		req, _ := http.NewRequest("GET", "127.0.0.1", nil)
		filters := kubernetes.FilterApplication(req)

		if !reflect.DeepEqual(filters, expectedFilters) {
			t.Fatalf("unexpected default queries values, got %v expected %v", expectedFilters, filters)
		}

	})

	t.Run("custom filters", func(t *testing.T) {
		expectedFilters := kubernetes.FilterApplications{
			Offset:        20,
			Limit:         40,
			Clusters:      []string{"cluster", "cluster1"},
			Namespaces:    []string{"namespace", "namespace1"},
			Statuses:      []string{""},
			SortBy:        "time",
			SortDirection: "desc",
			From:          0,
			To:            0,
			Distinct:      true,
		}
		req, _ := http.NewRequest("GET", fmt.Sprintf("127.0.0.1?offset=%d&limit=%d&cluster=%s&namespace=%s&distinct=%s", expectedFilters.Offset, expectedFilters.Limit, strings.Join(expectedFilters.Clusters, ","), strings.Join(expectedFilters.Namespaces, ","), strconv.FormatBool(expectedFilters.Distinct)), nil)
		filters := kubernetes.FilterApplication(req)
		if !reflect.DeepEqual(filters, expectedFilters) {
			t.Fatalf("unexpected default queries values, got %v expected %v", expectedFilters, filters)
		}

	})

}
